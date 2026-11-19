// JSON: marshal, unmarshal, struct tags, exported fields, encoder/decoder.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First    string   `json:"first"`
	Last     string   `json:"last"`
	Age      int      `json:"age"`
	Sayings  []string `json:"sayings,omitempty"`
	password string   // unexported: never appears in the JSON
}

func main() {
	people := []person{
		{First: "James", Last: "Bond", Age: 32, Sayings: []string{"Shaken, not stirred"}, password: "007"},
		{First: "Miss", Last: "Moneypenny", Age: 27},
	}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Marshal:")
	fmt.Println(string(bs))

	pretty, _ := json.MarshalIndent(people[0], "", "  ")
	fmt.Println("\nMarshalIndent:")
	fmt.Println(string(pretty))

	var back []person
	if err := json.Unmarshal(bs, &back); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nUnmarshal: %+v\n", back)

	fmt.Println("\nEncoder -> os.Stdout:")
	_ = json.NewEncoder(os.Stdout).Encode(people[1])

	fmt.Println("\nDecoder <- strings.Reader:")
	var p person
	_ = json.NewDecoder(strings.NewReader(`{"first":"Q","age":60}`)).Decode(&p)
	fmt.Printf("%+v\n", p)

	// Unknown shapes: decode into map[string]any.
	var anything map[string]any
	_ = json.Unmarshal([]byte(`{"n": 1, "tags": ["a", "b"], "ok": true}`), &anything)
	for k, v := range anything {
		fmt.Printf("%s: %v (%T)\n", k, v, v)
	}
}
