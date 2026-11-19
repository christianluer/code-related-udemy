# 11 · Application: JSON, sort & bcrypt

Course section: *Application*, hands-on exercises #1 to #5. Now that you know pointers and methods,
you can read and use the standard library. Docs live at [pkg.go.dev/std](https://pkg.go.dev/std).

## JSON

```go
type user struct {
	First string `json:"first"`        // struct tags rename fields in the JSON
	Age   int    `json:"age,omitempty"`
	pass  string                        // lowercase = unexported = INVISIBLE to encoding/json
}

bs, err := json.Marshal(users)        // Go value -> JSON []byte
err = json.Unmarshal(bs, &users)      // JSON -> Go value (pass a POINTER so it can fill it)

json.NewEncoder(os.Stdout).Encode(v)  // stream to any io.Writer
json.NewDecoder(r).Decode(&v)         // stream from any io.Reader
```

- The #1 JSON gotcha: **only exported (Capitalized) fields are marshaled**.
- [mholt.github.io/json-to-go](https://mholt.github.io/json-to-go/) turns a JSON sample into Go structs.
- Marshal/Unmarshal work on `[]byte` in memory, while Encoder/Decoder work on streams.

## sort

```go
sort.Ints(xi); sort.Strings(xs)                 // classic
slices.Sort(xi)                                 // modern generic
sort.Slice(people, func(i, j int) bool {        // custom order
	return people[i].Age < people[j].Age
})
slices.SortFunc(people, func(a, b person) int { // modern: return negative, 0 or positive
	return cmp.Compare(a.Last, b.Last)
})
```

You can also implement `sort.Interface` (`Len`, `Less`, `Swap`) on a named slice type, such as
`type ByAge []person`, and call `sort.Sort(ByAge(people))`.

## bcrypt

Never store plain-text passwords. `golang.org/x/crypto/bcrypt` is a deliberately **slow, salted** hash:

```go
hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
err = bcrypt.CompareHashAndPassword(hash, []byte(attempt)) // nil means match
```

The same password hashes differently every time because of a random salt, so you compare with
`CompareHashAndPassword`, never with `==`.

## Exercises

| Exercise | Course |
|---|---|
| `MarshalUsers` | HOE #1 (export your fields!) |
| `UnmarshalUsers` | HOE #2 |
| `EncodeUsers` | HOE #3 (json.NewEncoder) |
| `SortSayings` | HOE #4 |
| `SortByAge` / `SortByLast` | HOE #5 |
| `HashPassword` / `CheckPassword` | bcrypt |
