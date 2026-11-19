// Packages: this program depends on package puppy, which depends on package dog,
// mirroring the course's  learn-to-code-go ← puppy ← dog  example.
//
// Every folder is a package. Only Capitalized identifiers are visible outside it.
package main

import (
	"fmt"
	"runtime"

	"github.com/christianluer/code-related-udemy/lessons/03-tooling-and-modules/examples/packages/puppy"
)

func main() {
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.BigBark())
	fmt.Println("puppy.From10():", puppy.From10())
	// puppy.secret() // compile error: secret is unexported

	fmt.Printf("\nbuilt for %s/%s with %s\n", runtime.GOOS, runtime.GOARCH, runtime.Version())
	fmt.Println("try: GOOS=windows go build -o /tmp/demo.exe ./lessons/03-tooling-and-modules/examples/packages")
}
