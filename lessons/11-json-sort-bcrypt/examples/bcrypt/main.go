// bcrypt: store password hashes, never passwords.
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := "password123"

	h1, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	h2, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	fmt.Println("hash 1:", string(h1))
	fmt.Println("hash 2:", string(h2), "(different salt, different hash)")

	for _, attempt := range []string{"password123", "password1234"} {
		err := bcrypt.CompareHashAndPassword(h1, []byte(attempt))
		fmt.Printf("login with %q: ok=%v\n", attempt, err == nil)
	}
}
