# 03 · Programming fundamentals, the terminal & Go modules

Course sections: *Programming fundamentals for beginners*, *Your development environment*,
*Go mod and dependency management*, *Housekeeping*, *Packages*, *Cross compile*, HOE #14 to #22.

## Terminology

| term | meaning | example |
|---|---|---|
| declare | create a variable of a TYPE | `var x int` |
| assign | give a variable a value | `x = 42` |
| initialize | declare + assign | `var x int = 42` |
| identifier | the name of a variable, func, const... | `x`, `fName` |
| keyword | reserved word | `func`, `var`, `for`, `go`... (25 total) |
| operator / operand | in `2 + 2`, `+` is the operator and the `2`s are operands | |
| statement | smallest standalone instruction | `x := 2 + 7` |
| expression | something that evaluates to a value | `2 + 7`, `42` |
| scope | where an identifier exists and is accessible | keep it narrow |

`( )` parens, `{ }` curly braces, `[ ]` brackets.

## The terminal

```bash
pwd; ls -la; ls -lh       # where am I, what's here (incl. hidden, human sizes)
cd folder; cd ..          # move around
mkdir dir; touch f.txt    # create
cat f.txt; nano f.txt     # read / edit
chmod 755 f.txt           # permissions: owner/group/world, r=4 w=2 x=1
rm f.txt; rm -rf dir      # delete (careful!)
env; go env; go version   # environment
file f.txt; ulimit -a; ps aux | wc -l
```

`-rwxr-xr-x` reads as: `-` file (`d` = dir), then `rwx` owner, `r-x` group, `r-x` world.

## Git & GitHub

```bash
git status
git add --all
git commit -m "message"
git push
git tag v1.0.0 && git push origin --tags   # version your module
git show v1.0.0
```

Use SSH keys to authenticate with GitHub (search "github ssh authentication" for the current steps).

**Semantic versioning** `vMAJOR.MINOR.PATCH`: MAJOR for breaking changes, MINOR for new
backwards-compatible features, PATCH for backwards-compatible bug fixes. Tags come in two kinds:
lightweight (`git tag v1.4`) and annotated (`git tag -a v1.4 -m "msg"`).

## Running Go code

```bash
mkdir hello && cd hello
go mod init github.com/you/hello   # creates go.mod: the module's name and dependencies
go run main.go     # compile + run
go run .           # run the package in this folder
go build           # produce an executable here
go install         # put the executable in $GOPATH/bin (go env GOPATH)
GOOS=linux GOARCH=amd64 go build   # cross compile (GOOS=windows, darwin...)
go help; go help env
```

- **Folders are packages.** Every file in a folder has the same `package` name.
- **Capitalization = visibility.** `Visible` is exported outside the package. `notVisible` is not.
  Go says exported/unexported, not public/private.
- Imports have **file scope**. Package-level identifiers have **package scope** across files.

## Modules & dependency management

- A **module** is a collection of packages with a `go.mod` at its root. Name it like
  `github.com/username/repo`.
- **Direct** dependencies are imported by your code. **Indirect** dependencies are the dependencies of your dependencies.
- `go get github.com/GoesToEleven/puppy@latest` (or `@v1.2.0`, or `@<commit hash>`) adds/updates a dependency.
- `go mod tidy` adds missing and removes unused requirements and updates `go.sum`.
- `go.sum` holds **checksums** of every dependency so builds are reproducible and tamper-proof.
- Avoid spaghetti code: write modular code split into packages and funcs.
  See Russ Cox, *Our Software Dependency Problem*.

## Housekeeping: checksums, hashing, encryption, communication

- **Checksum / hash**: a fixed-size digest of any input. A tiny change in the input gives a completely
  different hash. `shasum -a 256 file`. Use SHA-256. MD5 and SHA-1 are broken.
  Passwords should use a slow hash like bcrypt (lesson 11).
- **Symmetric encryption** uses one shared key (AES): fast. **Asymmetric encryption** uses public/private keys
  (RSA, ECC): secure key exchange over insecure channels.
- Communication modes: **simplex** is one way (TV), **half-duplex** is both ways but one at a time
  (walkie-talkie), and **full-duplex** is both ways at once (phone).

## Exercises

| Exercise | Course |
|---|---|
| `Increment` uses a package-level variable | HOE #14: scope |
| `IsExported` applies the capitalization rule | Packages |
| `SHA256Hex` computes a checksum | Checksums, HOE #23 |
| `Bump` bumps a semantic version | Tag git commits with version |
| `Platform` reports GOOS/GOARCH | Cross compile, HOE #17 |

**Try at the terminal too** (HOE #15 to #22): run `go run ./lessons/03-tooling-and-modules/examples/packages`,
then `go build -o /tmp/pkgdemo ./lessons/03-tooling-and-modules/examples/packages && /tmp/pkgdemo`,
then `GOOS=windows go build -o /tmp/pkgdemo.exe ./lessons/03-tooling-and-modules/examples/packages && file /tmp/pkgdemo.exe`.
