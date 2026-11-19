package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
)

// Progress is persisted to .learn-progress.json at the repo root.
type Progress struct {
	Passing map[string][]string `json:"passing"` // lesson ID -> passing test names
	Quiz    map[string]int      `json:"quiz"`    // lesson ID -> best score
}

func newProgress() *Progress {
	return &Progress{Passing: map[string][]string{}, Quiz: map[string]int{}}
}

func (a *App) progressPath() string { return filepath.Join(a.Root, ".learn-progress.json") }

func loadProgress(path string) (*Progress, error) {
	p := newProgress()
	b, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return p, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, p); err != nil {
		return nil, err
	}
	if p.Passing == nil {
		p.Passing = map[string][]string{}
	}
	if p.Quiz == nil {
		p.Quiz = map[string]int{}
	}
	return p, nil
}

// Passed reports whether test last passed.
func (p *Progress) Passed(lesson, test string) bool {
	for _, t := range p.Passing[lesson] {
		if t == test {
			return true
		}
	}
	return false
}

// SetPassing replaces the passing tests for a lesson.
func (p *Progress) SetPassing(lesson string, tests []string) {
	sort.Strings(tests)
	p.Passing[lesson] = tests
}

// Save writes progress to disk.
func (p *Progress) Save(path string) error {
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}
