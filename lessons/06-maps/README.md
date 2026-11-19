# 06 · Grouping data: maps

Course section: *Grouping data values - map*, hands-on exercises #49 to #52.

A map is a **key/value store**: an unordered group of values of one type (the element type) indexed
by unique keys of another type (the key type). The key type must be **comparable** (`==` works on it):
strings, ints, structs of comparable fields, and so on. Slices can't be keys. Lookups are fast.

```go
m := map[string]int{          // map literal
	"James":      42,
	"Moneypenny": 32,
}
m["Q"] = 64                   // add / update
age := m["James"]             // access
fmt.Println(len(m))           // number of entries
delete(m, "Q")                // delete (no error if the key is missing)

for k, v := range m { }       // iteration order is RANDOM, on purpose

if v, ok := m["Bond"]; ok { } // comma ok: tell "missing" apart from "zero value"

counts := make(map[string]int) // make an empty, ready-to-use map
var nilMap map[string]int      // nil map: reads work, WRITES PANIC
```

- A missing key returns the **zero value** of the element type. Use comma ok when it matters.
- Need sorted output? Collect the keys, sort them, then range over the sorted keys
  (`slices.Sorted(maps.Keys(m))` in modern Go).
- Maps are reference-like: passing a map to a func lets the func modify the caller's map.
- `map[string][]string` is a perfect fit for "person → favorite things".

## Counting words in a book

```go
counts := map[string]int{}
for _, w := range strings.Fields(text) {
	counts[strings.ToLower(w)]++   // the zero value 0 makes this work for new words
}
```

## Exercises

| Exercise | Course |
|---|---|
| `Favorites` builds a map[string][]string | HOE #49 |
| `AddFavorite` appends to a value | HOE #50 |
| `Remove` uses delete + comma ok | HOE #51 |
| `WordFrequency` counts words | HOE #52, counting words in a book |
| `SortedKeys` gives deterministic order | Map - for range |
| `TopWords` gives the most frequent words | counting words in a book |
