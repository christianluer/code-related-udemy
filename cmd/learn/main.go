// Command learn is an interactive companion for the "Learn To Code Go" course.
//
// It lists the lessons, prints the notes, runs examples, checks your exercise
// solutions against the tests, quizzes you, and tracks your progress.
//
//	go run ./cmd/learn          # interactive menu
//	go run ./cmd/learn help     # all commands
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	root, err := findRoot()
	if err != nil {
		fatal(err)
	}
	app := &App{Root: root, In: bufio.NewReader(os.Stdin)}
	if err := app.Load(); err != nil {
		fatal(err)
	}

	if len(os.Args) > 1 {
		if err := app.Dispatch(os.Args[1:]); err != nil {
			fatal(err)
		}
		return
	}
	app.REPL()
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, red("error: ")+err.Error())
	os.Exit(1)
}

// Dispatch executes a single command.
func (a *App) Dispatch(args []string) error {
	cmd, rest := strings.ToLower(args[0]), args[1:]
	switch cmd {
	case "help", "h", "?":
		a.Help()
	case "list", "ls", "l":
		a.List()
	case "next", "n":
		return a.Next()
	case "read", "r", "notes":
		return a.withLesson(rest, a.Read)
	case "examples", "ex":
		return a.withLesson(rest, a.Examples)
	case "run":
		if len(rest) < 1 {
			return fmt.Errorf("usage: run <lesson> [example]")
		}
		l, err := a.Find(rest[0])
		if err != nil {
			return err
		}
		name := ""
		if len(rest) > 1 {
			name = rest[1]
		}
		return a.Run(l, name)
	case "check", "c", "test":
		return a.withLesson(rest, func(l *Lesson) error { _, err := a.Check(l, true); return err })
	case "watch", "w":
		return a.withLesson(rest, a.Watch)
	case "hint":
		return a.withLesson(rest, a.Hints)
	case "solution", "sol":
		return a.withLesson(rest, a.Solution)
	case "quiz", "q":
		return a.withLesson(rest, a.Quiz)
	case "verify":
		return a.Verify()
	case "reset":
		return a.Reset()
	default:
		return fmt.Errorf("unknown command %q (try \"help\")", cmd)
	}
	return nil
}

func (a *App) withLesson(args []string, fn func(*Lesson) error) error {
	var l *Lesson
	var err error
	if len(args) == 0 {
		l = a.firstIncomplete()
		if l == nil {
			return fmt.Errorf("all lessons complete, pass a lesson number")
		}
		fmt.Println(dim("(using next incomplete lesson: " + l.ID + ")"))
	} else if l, err = a.Find(args[0]); err != nil {
		return err
	}
	return fn(l)
}

// REPL runs the interactive menu.
func (a *App) REPL() {
	banner()
	a.List()
	fmt.Println()
	fmt.Println(dim("Type a command (\"help\" for the list, \"quit\" to exit)."))
	for {
		fmt.Print(cyan("learn> "))
		line, err := a.In.ReadString('\n')
		if err != nil {
			fmt.Println()
			return
		}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		if f := strings.ToLower(fields[0]); f == "quit" || f == "exit" {
			return
		}
		if err := a.Dispatch(fields); err != nil {
			fmt.Println(red("error: ") + err.Error())
		}
	}
}

// Help prints all commands.
func (a *App) Help() {
	fmt.Println(bold("Commands") + dim("  (<lesson> is a number like 4, or part of a name like \"slices\"; omit it to use the next incomplete lesson)"))
	rows := [][2]string{
		{"list", "show all lessons and your progress"},
		{"next", "show what to work on next"},
		{"read <lesson>", "print the lesson notes"},
		{"examples <lesson>", "list runnable examples"},
		{"run <lesson> [example]", "run an example (all of them if omitted)"},
		{"check <lesson>", "run the exercise tests and record progress"},
		{"watch <lesson>", "re-run the tests every time you save"},
		{"hint <lesson>", "show hints for the exercises you haven't solved"},
		{"solution <lesson>", "show the reference solution"},
		{"quiz <lesson>", "take the lesson quiz"},
		{"verify", "check that every reference solution passes"},
		{"reset", "wipe your progress"},
		{"quit", "leave the interactive menu"},
	}
	for _, r := range rows {
		fmt.Printf("  %-26s %s\n", cyan(r[0]), r[1])
	}
}

func banner() {
	fmt.Println(cyan(`
   ʕ◔ϖ◔ʔ   Learn To Code Go`))
	fmt.Println(dim("           interactive companion to Todd McLeod's course"))
	fmt.Println()
}
