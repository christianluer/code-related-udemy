//go:build solution

package exercises

import (
	"encoding/json"
	"io"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

func MarshalUsers(users []User) (string, error) {
	bs, err := json.Marshal(users)
	return string(bs), err
}

func UnmarshalUsers(data string) ([]User, error) {
	var users []User
	err := json.Unmarshal([]byte(data), &users)
	return users, err
}

func EncodeUsers(w io.Writer, users []User) error {
	return json.NewEncoder(w).Encode(users)
}

func SortSayings(users []User) {
	for _, u := range users {
		sort.Strings(u.Sayings)
	}
}

func SortByAge(users []User) {
	sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })
}

func SortByLast(users []User) {
	sort.Slice(users, func(i, j int) bool {
		if users[i].Last != users[j].Last {
			return users[i].Last < users[j].Last
		}
		return users[i].First < users[j].First
	})
}

func HashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(h), err
}

func CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
