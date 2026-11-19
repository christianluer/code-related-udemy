// Interpreted ("...") vs raw (`...`) string literals.
package main

import "fmt"

func main() {
	interpreted := "Tab:\tNewline:\nQuote:\" Backslash:\\"
	raw := `Tab:\tNewline:\nQuote:" Backslash:\
...and raw literals can span
multiple lines.`

	fmt.Println("interpreted:")
	fmt.Println(interpreted)
	fmt.Println("\nraw:")
	fmt.Println(raw)

	// %q shows the escaped form of a string, handy for debugging.
	fmt.Printf("\n%%q of interpreted: %q\n", interpreted)
	fmt.Printf("%%q of raw:         %q\n", raw)

	// Raw literals are perfect for regexes, Windows paths and JSON.
	fmt.Println(`C:\Users\gopher\go`)
	fmt.Println(`{"name": "gopher", "age": 14}`)
}
