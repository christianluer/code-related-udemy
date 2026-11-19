# 07 · Grouping data: structs

Course section: *Grouping data values - structs*, hands-on exercises #53 to #56.

A struct is a **composite** (aggregate, complex) data type: it composes values of *different*
types into one data structure.

```go
type person struct {
	first   string
	last    string
	age     int
	flavors []string
}

p1 := person{first: "James", last: "Bond", age: 32}   // composite literal with field names
p2 := person{"Miss", "Moneypenny", 27, nil}           // positional (brittle, avoid)
fmt.Println(p1.first)                                 // dot notation
fmt.Printf("%+v\n", p1)                               // %+v prints field names
```

## Embedded structs & promotion

```go
type secretAgent struct {
	person          // embedded: no field name
	licenseToKill bool
}
sa := secretAgent{person: person{first: "James"}, licenseToKill: true}
sa.first          // the inner type's fields (and methods) are PROMOTED to the outer type
sa.person.first   // still reachable explicitly
```

If the outer type declares a field or method with the same name, it **overrides** (shadows) the promoted one.

## Anonymous structs

```go
p := struct {
	first   string
	friends map[string]int
}{
	first:   "James",
	friends: map[string]int{"Moneypenny": 555},
}
```

These are good for one-off data such as test tables or decoding a JSON response.

## Is Go object oriented? It's all about TYPE

| OOP idea | Go |
|---|---|
| class | you create a **TYPE** (`type person struct{...}`) |
| instance / object | you create a **VALUE** of a TYPE |
| encapsulation | fields (state) + methods (behavior), exported vs unexported |
| inheritance / reuse | **embedded types** (composition, not inheritance) |
| polymorphism | **interfaces** (lesson 08) |
| overriding | promotion + shadowing |

- **Why create a type?** To attach methods, and to self-document code (`type Duration int64` in package `time`).
- **Named vs anonymous types**: you can assign an anonymous type to a named type with the same
  underlying type, but you cannot assign one named type to a different named type without a conversion.
- **Padding & alignment**: fields are aligned in memory. Organize fields logically for readability first,
  and only order them largest to smallest when performance demands it.

## Exercises

| Exercise | Course |
|---|---|
| `NewPerson` builds a struct with a slice | HOE #53 |
| `ByLastName` builds a map of structs | HOE #54 |
| `NewVehicle` / `IsElectric` use an embedded struct | HOE #55 |
| `NewProfile` uses an anonymous struct | HOE #56 |
| `FullNames` ranges over structs | HOE #53 |
