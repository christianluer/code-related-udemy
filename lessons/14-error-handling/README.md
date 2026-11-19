# 14 · Error handling

Course section: *Error handling*, hands-on exercises #1 to #5.

Go has no try/catch/finally (see the [FAQ](https://go.dev/doc/faq#exceptions)). Errors are ordinary
**values** returned from functions:

> Error values in Go aren't special, they are just values like any other, and so you have the
> entire language at your disposal. (Rob Pike)

```go
type error interface { Error() string }  // the built-in interface

f, err := os.Open("file.txt")
if err != nil {           // check errors. Always, always, always. (*almost always)
	return fmt.Errorf("open config: %w", err)
}
defer f.Close()
```

"Write the code with errors before you write the code without errors."

## Printing & logging

| call | effect |
|---|---|
| `fmt.Println(err)` | print to stdout |
| `log.Println(err)` | print with a timestamp (use `log.SetOutput(f)` to log to a file) |
| `log.Fatalln(err)` | log, then `os.Exit(1)`. **Deferred funcs do NOT run.** |
| `log.Panicln(err)` | log, then `panic`. Deferred funcs run, and you can `recover` |
| `panic(err)` | stop normal execution and unwind the stack |

## defer, panic & recover

```go
func safe() (err error) {
	defer func() {
		if r := recover(); r != nil {      // recover only works inside a deferred func
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	riskyThing()
	return nil
}
```

Reserve panic for truly exceptional, unrecoverable situations, such as programmer errors.

## Errors with info

```go
errors.New("something broke")                  // a simple error value
fmt.Errorf("user %d: %w", id, err)             // add context; %w WRAPS err
var ErrNotFound = errors.New("not found")      // sentinel error (package-level)

errors.Is(err, ErrNotFound)                    // is ErrNotFound anywhere in the chain?
var pe *fs.PathError
errors.As(err, &pe)                            // is there a *fs.PathError in the chain? Fill pe
errors.Join(err1, err2)                        // combine several errors (Go 1.20+)
```

Custom error types are structs with an `Error() string` method. They can carry fields, like the
course's `sqrtError{lat, long string; err error}`.

## Exercises

| Exercise | Course |
|---|---|
| `ReadConfig` checks and wraps an error | HOE #1, #2 |
| `ParseAge` uses sentinel errors + `%w` | HOE #2 |
| `CustomErr` / `Describe` custom type | HOE #3 |
| `Sqrt` + `SqrtError` with lat/long | HOE #4 |
| `SafeDivide` panic + recover | Recover |
| `ValidateAll` errors.Join | Errors with info |
