package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Run executes one example, or every example in the lesson.
func (a *App) Run(l *Lesson, name string) error {
	names := a.examples(l)
	if name != "" {
		found := false
		for _, n := range names {
			if n == name || strings.Contains(n, name) {
				names, found = []string{n}, true
				break
			}
		}
		if !found {
			return fmt.Errorf("no example %q in %s (try \"examples %d\")", name, l.ID, l.Number)
		}
	}
	for _, n := range names {
		pkg := "./" + rel(a.Root, filepath.Join(l.Dir, "examples", n))
		fmt.Println(bold(yellow("▶ go run " + pkg)))
		cmd := exec.Command("go", "run", pkg)
		cmd.Dir = a.Root
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(red("example exited with: " + err.Error()))
		}
		fmt.Println()
	}
	return nil
}

type testEvent struct {
	Action string
	Test   string
	Output string
}

type testResult struct {
	passed bool
	output []string
}

// Check runs the lesson's exercise tests, prints a report, and saves progress.
// It returns true when every exercise passes.
func (a *App) Check(l *Lesson, verbose bool) (bool, error) {
	pkg := "./" + rel(a.Root, filepath.Join(l.Dir, "exercises"))
	cmd := exec.Command("go", "test", "-json", "-race", "-count=1", "-timeout=60s", pkg)
	cmd.Dir = a.Root
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, runErr := cmd.Output()

	results := map[string]*testResult{}
	var pkgOutput []string
	sc := bufio.NewScanner(bytes.NewReader(out))
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		var ev testEvent
		if err := json.Unmarshal(sc.Bytes(), &ev); err != nil {
			pkgOutput = append(pkgOutput, sc.Text())
			continue
		}
		top := strings.SplitN(ev.Test, "/", 2)[0]
		if top == "" {
			if ev.Output != "" {
				pkgOutput = append(pkgOutput, strings.TrimRight(ev.Output, "\n"))
			}
			continue
		}
		r := results[top]
		if r == nil {
			r = &testResult{}
			results[top] = r
		}
		switch {
		case ev.Action == "output" && keepLine(ev.Output):
			r.output = append(r.output, strings.TrimRight(ev.Output, "\n"))
		case ev.Action == "pass" && ev.Test == top:
			r.passed = true
		}
	}

	fmt.Printf("%s %s\n", bold("Checking"), l.Meta.Title)
	if len(results) == 0 {
		fmt.Println(red("✗ the code does not compile or the tests could not run:"))
		for _, line := range append(pkgOutput, strings.Split(strings.TrimSpace(stderr.String()), "\n")...) {
			if strings.TrimSpace(line) != "" {
				fmt.Println("    " + line)
			}
		}
		return false, nil
	}

	var passing []string
	for _, t := range l.TestList {
		r, ok := results[t]
		switch {
		case !ok:
			fmt.Printf("  %s %s %s\n", yellow("?"), t, dim("(did not run: an earlier test probably panicked)"))
		case r.passed:
			passing = append(passing, t)
			fmt.Printf("  %s %s\n", green("✓"), t)
		default:
			fmt.Printf("  %s %s\n", red("✗"), t)
			if verbose {
				for _, line := range r.output {
					fmt.Println("      " + dim(strings.TrimSpace(line)))
				}
			}
		}
	}
	if runErr != nil && len(passing) == len(l.TestList) {
		for _, line := range pkgOutput {
			fmt.Println("    " + line)
		}
	}

	a.Progress.SetPassing(l.ID, passing)
	if err := a.Progress.Save(a.progressPath()); err != nil {
		return false, err
	}

	done := len(passing) == len(l.TestList)
	fmt.Printf("\n  %s %s\n", bar(len(passing), len(l.TestList)), dim(fmt.Sprintf("%d/%d passing", len(passing), len(l.TestList))))
	if done {
		fmt.Println(green("  🎉 Lesson complete! ") + dim(fmt.Sprintf("Try \"quiz %d\" or move on with \"next\".", l.Number)))
	} else {
		fmt.Println(dim(fmt.Sprintf("  Stuck? \"hint %d\" shows hints for the failing exercises.", l.Number)))
	}
	return done, nil
}

func keepLine(s string) bool {
	t := strings.TrimSpace(s)
	for _, p := range []string{"=== RUN", "=== PAUSE", "=== CONT", "--- PASS", "--- FAIL", "PASS", "FAIL"} {
		if strings.HasPrefix(t, p) {
			return false
		}
	}
	return t != ""
}

// Watch re-runs Check whenever a .go file in the exercises folder changes.
func (a *App) Watch(l *Lesson) error {
	dir := filepath.Join(l.Dir, "exercises")
	last := time.Time{}
	fmt.Println(dim("Watching " + rel(a.Root, dir) + " (Ctrl+C to stop)"))
	for {
		latest := latestMod(dir)
		if latest.After(last) {
			last = latest
			fmt.Print("\033[H\033[2J")
			fmt.Println(dim("Watching " + rel(a.Root, dir) + " (Ctrl+C to stop). Save a file to re-run."))
			done, err := a.Check(l, true)
			if err != nil {
				return err
			}
			if done {
				return nil
			}
		}
		time.Sleep(700 * time.Millisecond)
	}
}

func latestMod(dir string) time.Time {
	var latest time.Time
	_ = filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".go") {
			return nil
		}
		if info, err := d.Info(); err == nil && info.ModTime().After(latest) {
			latest = info.ModTime()
		}
		return nil
	})
	return latest
}

// Verify runs every exercise test against the reference solutions.
func (a *App) Verify() error {
	fmt.Println(bold("Running every test against the reference solutions (-tags solution)"))
	cmd := exec.Command("go", "test", "-tags", "solution", "-race", "-count=1", "./lessons/...")
	cmd.Dir = a.Root
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}
