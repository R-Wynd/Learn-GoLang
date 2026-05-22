package main 
import "fmt"

func main() {

	// Six main points of GO
	// 1. Go is statically typed language
	// 2. Strongly typed language
	// 3. Go is complied
	// 4. Fast Complilation time
	// 5. Build in concurrency 
	// 6. Simplicity 

	// 1. Go is statically typed language : The compiler knows the type of every variable before the program runs. Many bugs are caught early, before you deploy.

	// 2. Go is strongly typed language : You cannot mix incompatible data types directly. 
	//    Example: adding a string and integer causes a compile-time error. 
	//    This improves type safety and reduces unexpected behavior.

	// 3. Go is compiled : Go source code is converted directly into machine code 
	//    before execution. The generated binary can run without needing an interpreter.

	// 4. Fast compilation time : Go is designed to compile very quickly even for 
	//    large projects. This improves developer productivity and reduces build times.

	// 5. Built-in concurrency : Go provides goroutines and channels to handle multiple 
	//    tasks concurrently with very little code. This makes it easy to build scalable 
	//    and high-performance applications.

	// 6. Simplicity : Go has a small syntax, minimal keywords, and straightforward 
	//    language design. It is easy to read, learn, maintain, and debug.

	// Why Go is widely used in the DevOps world:

	// 1. Single binary deployment : Go applications compile into a standalone binary.
	//    Easy to deploy in servers, containers, and Kubernetes clusters without extra runtimes.

	// 2. Fast performance : Go provides performance close to C/C++ while remaining simpler to write and maintain.

	// 3. Fast compilation : Large infrastructure projects compile quickly, improving developer productivity.

	// 4. Built-in concurrency : Goroutines and channels make it easy to handle parallel tasks like monitoring,
	//    networking, controllers, and distributed systems.

	// 5. Excellent networking support : Go has strong built-in libraries for HTTP, TCP, DNS, TLS, and APIs,
	//    making it ideal for cloud and infrastructure tools.

	// 6. Simple and maintainable : Go code is clean, readable, and easier for teams to maintain at scale.

	// 7. Cross-platform support : Go can easily build binaries for Linux, Windows, macOS, ARM, and x86.

	// 8. Minimal dependencies : Go’s rich standard library reduces external dependencies and operational complexity.

	// 9. Kubernetes ecosystem : Kubernetes and many cloud-native tools are written in Go,
	//    making Go the default language for DevOps and platform engineering.

	// 10. Strong CLI ecosystem : Go is excellent for building fast and portable command-line tools.

	// Examples of DevOps tools written in Go:
	// Docker, Kubernetes, Terraform, Helm, Prometheus, Vault, Consul, etc.


	// Simple Print statement in Go
	fmt.Println("Hello World!!")
}