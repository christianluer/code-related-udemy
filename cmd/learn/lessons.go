package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// App holds the state of the CLI.
type App struct {
	Root     string
	In       *bufio.Reader
	Lessons  []*Lesson
	Progress *Progress
}

// Lesson is one directory under lessons/.
type Lesson struct {
	ID       string // directory name, e.g. "04-control-flow"
	Number   int
	Dir      string
	Meta     Meta
	TestList []string // top-level test names in exercises/exercise_test.go
}

// Meta is the content of lesson.json.
type Meta struct {
	Title   string            `json:"title"`
	Section string            `json:"course_section"`
	Hints   map[string]string `json:"hints"`
	Quiz    []Question        `json:"quiz"`
}

// Question is a multiple choice quiz question.
type Question struct {
	Q       string   `json:"q"`
	Options []string `json:"options"`
	Answer  int      `json:"answer"`
	Why     string   `json:"why"`
}

var testFuncRe = regexp.MustCompile(`(?m)^func (Test\w+\(t \*testing\.T\)|Example\w*\(\))`)

func findRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "lessons")); err == nil {
			if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
				return dir, nil
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("could not find the course repo (run this from inside it)")
		}
		dir = parent
	}
}

// Load discovers lessons and reads saved progress.
func (a *App) Load() error {
	entries, err := os.ReadDir(filepath.Join(a.Root, "lessons"))
	if err != nil {
		return err
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		num, err := strconv.Atoi(strings.SplitN(e.Name(), "-", 2)[0])
		if err != nil {
			continue
		}
		l := &Lesson{ID: e.Name(), Number: num, Dir: filepath.Join(a.Root, "lessons", e.Name())}
		if b, err := os.ReadFile(filepath.Join(l.Dir, "lesson.json")); err == nil {
			if err := json.Unmarshal(b, &l.Meta); err != nil {
				return fmt.Errorf("%s/lesson.json: %w", l.ID, err)
			}
		}
		if l.Meta.Title == "" {
			l.Meta.Title = l.ID
		}
		if b, err := os.ReadFile(filepath.Join(l.Dir, "exercises", "exercise_test.go")); err == nil {
			for _, m := range testFuncRe.FindAllSubmatch(b, -1) {
				l.TestList = append(l.TestList, strings.SplitN(string(m[1]), "(", 2)[0])
			}
		}
		a.Lessons = append(a.Lessons, l)
	}
	sort.Slice(a.Lessons, func(i, j int) bool { return a.Lessons[i].Number < a.Lessons[j].Number })

	a.Progress, err = loadProgress(a.progressPath())
	return err
}

// Find resolves a lesson by number or by a substring of its name.
func (a *App) Find(key string) (*Lesson, error) {
	if n, err := strconv.Atoi(key); err == nil {
		for _, l := range a.Lessons {
			if l.Number == n {
				return l, nil
			}
		}
		return nil, fmt.Errorf("no lesson number %d", n)
	}
	key = strings.ToLower(key)
	for _, l := range a.Lessons {
		if strings.Contains(l.ID, key) || strings.Contains(strings.ToLower(l.Meta.Title), key) {
			return l, nil
		}
	}
	return nil, fmt.Errorf("no lesson matching %q", key)
}

func (a *App) solved(l *Lesson) int {
	n := 0
	for _, t := range l.TestList {
		if a.Progress.Passed(l.ID, t) {
			n++
		}
	}
	return n
}

func (a *App) complete(l *Lesson) bool {
	return len(l.TestList) > 0 && a.solved(l) == len(l.TestList)
}

func (a *App) firstIncomplete() *Lesson {
	for _, l := range a.Lessons {
		if !a.complete(l) {
			return l
		}
	}
	return nil
}

// List prints the lessons with progress bars.
func (a *App) List() {
	totalDone, total := 0, 0
	for _, l := range a.Lessons {
		done, n := a.solved(l), len(l.TestList)
		totalDone += done
		total += n
		mark := "  "
		switch {
		case a.complete(l):
			mark = green("✓ ")
		case done > 0:
			mark = yellow("… ")
		}
		quiz := ""
		if s, ok := a.Progress.Quiz[l.ID]; ok {
			quiz = dim(fmt.Sprintf("  quiz %d/%d", s, len(l.Meta.Quiz)))
		}
		fmt.Printf("%s%2d  %-38s %s %s%s\n", mark, l.Number, l.Meta.Title, bar(done, n), dim(fmt.Sprintf("%d/%d", done, n)), quiz)
	}
	fmt.Printf("\n    %s %s\n", bold("Overall"), bar(totalDone, total)+dim(fmt.Sprintf(" %d/%d exercises", totalDone, total)))
}

func bar(done, total int) string {
	const width = 12
	if total == 0 {
		return dim(strings.Repeat("·", width))
	}
	filled := done * width / total
	return green(strings.Repeat("█", filled)) + dim(strings.Repeat("░", width-filled))
}

// Next suggests what to do next.
func (a *App) Next() error {
	l := a.firstIncomplete()
	if l == nil {
		fmt.Println(green("🎉 Every exercise passes. You finished the course!"))
		return nil
	}
	fmt.Printf("Next up: %s %s\n", bold(fmt.Sprintf("%d. %s", l.Number, l.Meta.Title)), dim("("+l.Meta.Section+")"))
	fmt.Printf("  1. read the notes        %s\n", cyan(fmt.Sprintf("read %d", l.Number)))
	fmt.Printf("  2. run the examples      %s\n", cyan(fmt.Sprintf("run %d", l.Number)))
	fmt.Printf("  3. edit                  %s\n", cyan(rel(a.Root, filepath.Join(l.Dir, "exercises", "exercise.go"))))
	fmt.Printf("  4. check as you go       %s\n", cyan(fmt.Sprintf("watch %d", l.Number)))
	fmt.Printf("  5. test yourself         %s\n", cyan(fmt.Sprintf("quiz %d", l.Number)))
	return nil
}

// Read prints the lesson notes.
func (a *App) Read(l *Lesson) error {
	b, err := os.ReadFile(filepath.Join(l.Dir, "README.md"))
	if err != nil {
		return err
	}
	fmt.Println(renderMarkdown(string(b)))
	return nil
}

func (a *App) examples(l *Lesson) []string {
	entries, _ := os.ReadDir(filepath.Join(l.Dir, "examples"))
	var names []string
	for _, e := range entries {
		if e.IsDir() {
			names = append(names, e.Name())
		}
	}
	return names
}

// Examples lists the runnable examples for a lesson.
func (a *App) Examples(l *Lesson) error {
	names := a.examples(l)
	if len(names) == 0 {
		fmt.Println(dim("no examples in this lesson"))
		return nil
	}
	fmt.Println(bold("Examples in " + l.Meta.Title))
	for _, n := range names {
		fmt.Printf("  %-28s %s\n", n, dim(fmt.Sprintf("run %d %s", l.Number, n)))
	}
	return nil
}

// Hints shows hints for unsolved exercises.
func (a *App) Hints(l *Lesson) error {
	shown := 0
	for _, t := range l.TestList {
		if a.Progress.Passed(l.ID, t) {
			continue
		}
		h, ok := l.Meta.Hints[t]
		if !ok {
			continue
		}
		fmt.Printf("%s %s\n   %s\n", yellow("💡"), bold(t), h)
		shown++
	}
	if shown == 0 {
		fmt.Println(dim("no hints left: either everything passes, or run \"check\" first"))
	}
	return nil
}

// Solution prints the reference solution after confirmation.
func (a *App) Solution(l *Lesson) error {
	fmt.Print(yellow("Peeking is fine, but try first! Show the solution? [y/N] "))
	ans, _ := a.In.ReadString('\n')
	if !strings.HasPrefix(strings.ToLower(strings.TrimSpace(ans)), "y") {
		return nil
	}
	b, err := os.ReadFile(filepath.Join(l.Dir, "exercises", "solution.go"))
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

// Quiz runs the interactive multiple choice quiz.
func (a *App) Quiz(l *Lesson) error {
	if len(l.Meta.Quiz) == 0 {
		fmt.Println(dim("no quiz for this lesson"))
		return nil
	}
	fmt.Println(bold("Quiz: " + l.Meta.Title))
	score := 0
	for i, q := range l.Meta.Quiz {
		fmt.Printf("\n%s %s\n", cyan(fmt.Sprintf("Q%d.", i+1)), q.Q)
		for j, o := range q.Options {
			fmt.Printf("   %c) %s\n", 'a'+j, o)
		}
		var pick int
		for {
			fmt.Print("   your answer: ")
			line, err := a.In.ReadString('\n')
			if err != nil {
				return nil
			}
			line = strings.ToLower(strings.TrimSpace(line))
			if len(line) == 1 && line[0] >= 'a' && int(line[0]-'a') < len(q.Options) {
				pick = int(line[0] - 'a')
				break
			}
			fmt.Println(dim("   type a letter"))
		}
		if pick == q.Answer {
			score++
			fmt.Println("   " + green("✓ correct"))
		} else {
			fmt.Printf("   %s the answer is %c) %s\n", red("✗"), 'a'+q.Answer, q.Options[q.Answer])
		}
		if q.Why != "" {
			fmt.Println("   " + dim(q.Why))
		}
	}
	fmt.Printf("\nScore: %s\n", bold(fmt.Sprintf("%d/%d", score, len(l.Meta.Quiz))))
	if score > a.Progress.Quiz[l.ID] {
		a.Progress.Quiz[l.ID] = score
		return a.Progress.Save(a.progressPath())
	}
	return nil
}

// Reset wipes progress.
func (a *App) Reset() error {
	fmt.Print(yellow("Delete all progress? [y/N] "))
	ans, _ := a.In.ReadString('\n')
	if !strings.HasPrefix(strings.ToLower(strings.TrimSpace(ans)), "y") {
		return nil
	}
	a.Progress = newProgress()
	return a.Progress.Save(a.progressPath())
}

func rel(root, p string) string {
	if r, err := filepath.Rel(root, p); err == nil {
		return r
	}
	return p
}

// renderMarkdown adds light terminal styling to headings and code fences.
func renderMarkdown(s string) string {
	var out strings.Builder
	inCode := false
	for _, line := range strings.Split(s, "\n") {
		switch {
		case strings.HasPrefix(line, "```"):
			inCode = !inCode
			continue
		case inCode:
			out.WriteString("    " + cyan(line))
		case strings.HasPrefix(line, "#"):
			out.WriteString(bold(yellow(strings.TrimLeft(line, "# "))))
		default:
			out.WriteString(line)
		}
		out.WriteByte('\n')
	}
	return out.String()
}
