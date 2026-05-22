# Learn Go: Project Structure
## Go Learning Examples

cmd/ : This folder holds **one runnable program per topic**, instead of putting everything in a single root `main.go`.

## Why split by topic?

- **Focus** — each folder gives one idea (variables, datatypes, pointers, etc.)
- **Easier to run** — run only what you're studying
- **Less clutter** — root `main.go` doesn't grow into one long file
- **Common Go layout** — `cmd/` is the usual place for small executables (convention only; Go doesn't treat the name specially)

## Structure

Each subfolder is its own `package main` with a `main()` function


## How to run
From the project root (`Learn-GoLang/`):
```bash
go run ./cmd/1_variables or go run ./cmd/1_variables/main.go
go run ./cmd/2_datatypes