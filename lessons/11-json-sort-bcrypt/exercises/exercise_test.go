package exercises

import (
	"bytes"
	"encoding/json"
	"slices"
	"strings"
	"testing"
)

func agents() []User {
	return []User{
		{First: "James", Last: "Bond", Age: 32, Sayings: []string{"Shaken, not stirred", "Youth is no guarantee of innovation", "In his majesty's royal service"}},
		{First: "Miss", Last: "Moneypenny", Age: 27, Sayings: []string{"James, it is soo good to see you", "Would you like me to take care of that for you, James?", "I would really prefer to be a secret agent myself."}},
		{First: "M", Last: "Hmmmm", Age: 54, Sayings: []string{"Oh, James. You didn't.", "Dear God, what has James done now?", "Can someone please tell me where James Bond is?"}},
		{First: "Jenny", Last: "Bond", Age: 20},
	}
}

func names(us []User) []string {
	var out []string
	for _, u := range us {
		out = append(out, u.First)
	}
	return out
}

func TestMarshalUsers(t *testing.T) {
	s, err := MarshalUsers(agents()[:1])
	if err != nil {
		t.Fatal(err)
	}
	want := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]}]`
	if s != want {
		t.Errorf("MarshalUsers =\n%s\nwant\n%s", s, want)
	}
}

func TestUnmarshalUsers(t *testing.T) {
	data := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred"]},{"First":"M","Last":"Hmmmm","Age":54}]`
	us, err := UnmarshalUsers(data)
	if err != nil {
		t.Fatal(err)
	}
	if len(us) != 2 || us[0].First != "James" || us[1].Age != 54 || len(us[0].Sayings) != 1 {
		t.Errorf("UnmarshalUsers = %+v", us)
	}
	if _, err := UnmarshalUsers(`{not json`); err == nil {
		t.Error("UnmarshalUsers should return an error for invalid JSON")
	}
}

func TestEncodeUsers(t *testing.T) {
	var b bytes.Buffer
	if err := EncodeUsers(&b, agents()); err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(b.String(), "\n") {
		t.Error("expected the output of json.Encoder (it ends with a newline)")
	}
	var back []User
	if err := json.Unmarshal(b.Bytes(), &back); err != nil || len(back) != 4 {
		t.Errorf("encoded output did not round-trip: %v", err)
	}
}

func TestSortSayings(t *testing.T) {
	us := agents()
	SortSayings(us)
	for _, u := range us {
		if !slices.IsSorted(u.Sayings) {
			t.Errorf("%s's sayings are not sorted: %q", u.First, u.Sayings)
		}
	}
}

func TestSortByAge(t *testing.T) {
	us := agents()
	SortByAge(us)
	if got, want := names(us), []string{"Jenny", "Miss", "James", "M"}; !slices.Equal(got, want) {
		t.Errorf("SortByAge order = %v, want %v", got, want)
	}
}

func TestSortByLast(t *testing.T) {
	us := agents()
	SortByLast(us)
	if got, want := names(us), []string{"James", "Jenny", "M", "Miss"}; !slices.Equal(got, want) {
		t.Errorf("SortByLast order = %v, want %v (Bond, Bond, Hmmmm, Moneypenny)", got, want)
	}
}

func TestBcrypt(t *testing.T) {
	h, err := HashPassword("secret007")
	if err != nil {
		t.Fatal(err)
	}
	if h == "" || h == "secret007" || !strings.HasPrefix(h, "$2") {
		t.Fatalf("HashPassword returned %q, want a bcrypt hash starting with $2", h)
	}
	if !CheckPassword(h, "secret007") {
		t.Error("CheckPassword rejected the right password")
	}
	if CheckPassword(h, "secret008") {
		t.Error("CheckPassword accepted the wrong password")
	}
}
