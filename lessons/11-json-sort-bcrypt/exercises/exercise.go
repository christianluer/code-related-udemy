//go:build !solution

// Package exercises contains the exercises for lesson 11.
//
// Replace each TODO, then run:  go run ./cmd/learn check 11
package exercises

import "io"

// MarshalUsers returns users as a JSON string. (HOE #1)
func MarshalUsers(users []User) (string, error) {
	// TODO: encoding/json
	return "", nil
}

// UnmarshalUsers parses a JSON array of users. (HOE #2)
func UnmarshalUsers(data string) ([]User, error) {
	// TODO
	return nil, nil
}

// EncodeUsers streams users as JSON to w using json.NewEncoder. (HOE #3)
func EncodeUsers(w io.Writer, users []User) error {
	// TODO
	return nil
}

// SortSayings sorts every user's Sayings alphabetically, in place. (HOE #4)
func SortSayings(users []User) {
	// TODO
}

// SortByAge sorts users youngest first, in place. (HOE #5)
func SortByAge(users []User) {
	// TODO
}

// SortByLast sorts users by Last name, then First name, in place. (HOE #5)
func SortByLast(users []User) {
	// TODO
}

// HashPassword returns a bcrypt hash of password.
// Use bcrypt.MinCost to keep the tests fast (use DefaultCost in real code).
func HashPassword(password string) (string, error) {
	// TODO: golang.org/x/crypto/bcrypt
	return "", nil
}

// CheckPassword reports whether password matches the bcrypt hash.
func CheckPassword(hash, password string) bool {
	// TODO
	return false
}
