# Learn Go

# 6 Main points of Go

## 1. Statically Typed Language

The compiler knows the type of every variable before the program runs.
Many bugs are caught early during compilation before deployment.

---

## 2. Strongly Typed Language

Go does not allow mixing incompatible data types directly.

Example:

```go id="6fjlwm"
var a int = 10
var b string = "hello"

fmt.Println(a + b) // Compile-time error
```

This improves type safety and reduces unexpected behavior.

---

## 3. Compiled Language

Go source code is compiled directly into machine code before execution.

The generated binary can run independently without needing:

* an interpreter
* a virtual machine
* external runtime dependencies

---

## 4. Fast Compilation Time

Go is designed for extremely fast compilation, even for large projects.

Benefits:

* faster development cycle
* quick testing
* improved developer productivity

---

## 5. Built-in Concurrency

Go provides lightweight concurrency using goroutines and channels

This makes it easy to build:

* scalable systems
* distributed applications
* high-performance services

---

## 6. Simplicity

Go has:

* small syntax
* minimal keywords
* straightforward design

This makes the language:

* easy to learn
* easy to read
* easy to maintain
* easy to debug


-------------------------------------------


# Why Go is Widely Used in the DevOps World

## 1. Single Binary Deployment

Go applications compile into a standalone binary, making deployment simple and portable without requiring external runtimes or dependencies.

## 2. Fast Performance

Go provides high performance close to C/C++ while remaining easier to write, debug, and maintain.

## 3. Built-in Concurrency

Go supports concurrency using goroutines and channels, making it ideal for handling distributed systems, monitoring, networking, and cloud workloads.

## 4. Excellent Networking Support

Go has powerful built-in networking libraries for HTTP, TCP, DNS, TLS, APIs, and microservices communication.

## 5. Simplicity and Maintainability

Go has clean syntax and minimal complexity, making large infrastructure codebases easier to maintain across teams.

## 6. Kubernetes and Cloud-Native Ecosystem

Many major DevOps and cloud-native tools are written in Go, including:

* Docker
* Kubernetes
* Terraform
* Helm
* Prometheus

Because Kubernetes itself is written in Go, it became the default language for modern DevOps and platform engineering.


# Go Learning Examples

cmd/ : This folder holds **one runnable program per topic**, instead of putting everything in a single root `main.go`.

## Why split by topic?

- **Focus** — each folder gives one idea (variables, datatypes, pointers, etc.)
- **Easier to run** — run only what you're studying
- **Less clutter** — root `main.go` doesn't grow into one long file
- **Common Go layout** — `cmd/` is the usual place for small executables (convention only; Go doesn't treat the name specially)


## How to run
From the project root (`Learn-GoLang/`):
```bash
go run ./cmd/1_variables or go run ./cmd/1_variables/main.go
go run ./cmd/2_datatypes