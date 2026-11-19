// io.Writer: write the same way to a file, a byte buffer or stdout.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type person struct{ first string }

// writeOut writes to ANY io.Writer.
func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte("Hello, " + p.first + "!\n"))
	return err
}

func main() {
	p := person{first: "Jenny"}

	path := filepath.Join(os.TempDir(), "gopher-output.txt")
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var b bytes.Buffer

	_ = p.writeOut(f)         // *os.File implements io.Writer
	_ = p.writeOut(&b)        // *bytes.Buffer implements io.Writer
	_ = p.writeOut(os.Stdout) // so does os.Stdout

	fmt.Print("buffer holds: ", b.String())
	fmt.Println("file written to:", path)

	// bytes.Buffer basics
	buf := bytes.NewBufferString("Hello, ")
	buf.WriteString("World!")
	fmt.Println(buf.String())
	buf.Reset()
	buf.WriteString("Go programming language!")
	fmt.Println(buf.String())

	// string <-> []byte
	bs := []byte("Hi")
	fmt.Println(bs, string([]byte{72, 105}))

	// fmt.Fprintf writes formatted text to any io.Writer.
	fmt.Fprintf(os.Stdout, "%s has %d letters\n", p.first, len(p.first))
}
