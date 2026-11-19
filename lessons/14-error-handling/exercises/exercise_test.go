package exercises

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestReadConfig(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "app.conf")
	if err := os.WriteFile(good, []byte("debug=true"), 0o644); err != nil {
		t.Fatal(err)
	}
	if b, err := ReadConfig(good); err != nil || string(b) != "debug=true" {
		t.Errorf("ReadConfig(existing) = (%q, %v)", b, err)
	}

	missing := filepath.Join(dir, "missing.conf")
	_, err := ReadConfig(missing)
	if err == nil {
		t.Fatal("ReadConfig(missing) returned a nil error")
	}
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("error %q doesn't wrap fs.ErrNotExist (use %%w)", err)
	}
	if !strings.Contains(err.Error(), "missing.conf") {
		t.Errorf("error %q should mention the path", err)
	}
}

func TestParseAge(t *testing.T) {
	if n, err := ParseAge("42"); n != 42 || err != nil {
		t.Errorf("ParseAge(\"42\") = (%d, %v)", n, err)
	}
	for _, s := range []string{"-1", "151", "9000"} {
		if _, err := ParseAge(s); !errors.Is(err, ErrInvalidAge) {
			t.Errorf("ParseAge(%q) error = %v, want it to wrap ErrInvalidAge", s, err)
		}
	}
	_, err := ParseAge("forty-two")
	var ne *strconv.NumError
	if !errors.As(err, &ne) {
		t.Errorf("ParseAge(\"forty-two\") error = %v, want it to wrap a *strconv.NumError", err)
	}
}

func TestCustomErr(t *testing.T) {
	var err error = CustomErr{Info: "need more coffee"}
	if got := Describe(err); got != "custom error: need more coffee" {
		t.Errorf("Describe = %q", got)
	}
	if got := Describe(errors.New("plain")); got != "plain" {
		t.Errorf("Describe(errors.New(\"plain\")) = %q", got)
	}
}

func TestSqrt(t *testing.T) {
	if v, err := Sqrt(16); v != 4 || err != nil {
		t.Errorf("Sqrt(16) = (%v, %v), want (4, nil)", v, err)
	}
	_, err := Sqrt(-10.23)
	var se *SqrtError
	if !errors.As(err, &se) {
		t.Fatalf("Sqrt(-10.23) error = %v, want a *SqrtError", err)
	}
	if se.Lat != "50.2289 N" || se.Long != "99.4656 W" || se.Err == nil {
		t.Errorf("SqrtError = %+v", se)
	}
	if !strings.HasPrefix(err.Error(), "math error: 50.2289 N 99.4656 W") {
		t.Errorf("Error() = %q", err.Error())
	}
	if errors.Unwrap(err) != se.Err {
		t.Error("Unwrap should return the inner Err")
	}
}

func TestSafeDivide(t *testing.T) {
	if r, err := SafeDivide(10, 2); r != 5 || err != nil {
		t.Errorf("SafeDivide(10, 2) = (%d, %v), want (5, nil)", r, err)
	}
	r, err := SafeDivide(1, 0)
	if err == nil {
		t.Fatalf("SafeDivide(1, 0) = (%d, nil), want an error", r)
	}
	if !strings.Contains(err.Error(), "divide by zero") {
		t.Errorf("error %q should include the recovered panic message", err)
	}
}

func TestValidateAll(t *testing.T) {
	notEmpty := func(s string) error {
		if s == "" {
			return errors.New("empty value")
		}
		return nil
	}
	if err := ValidateAll([]string{"a", "b"}, notEmpty); err != nil {
		t.Errorf("all valid: got %v, want nil", err)
	}
	sentinel := fmt.Errorf("too long")
	check := func(s string) error {
		if len(s) > 3 {
			return sentinel
		}
		return notEmpty(s)
	}
	err := ValidateAll([]string{"", "ok", "toolong"}, check)
	if err == nil || !errors.Is(err, sentinel) || !strings.Contains(err.Error(), "empty value") {
		t.Errorf("ValidateAll = %v, want both failures joined", err)
	}
}
