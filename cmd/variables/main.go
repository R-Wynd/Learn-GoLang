package main

import "fmt"

func main(){
	// Variable in Go
	const name = "Aravind"
	var age = 22
	fmt.Println("Hello", name)
	fmt.Println("Your age is", age)

	var a int = 22           // explicit type
	var b = "Aravind"        // type inferred
	c := "Rwind@21"          // short declaration (:=), type inferred
	fmt.Println(a,b,c)

	// Print prints values with no guaranteed spaces and no newline; 
	// Println prints values with spaces between them and adds a newline; 
	// Printf prints using a format string (%s, %d, etc.) and only adds a newline if you put \n in the format.



	// String formating
	// fmt.Printf with format verbs
	// %s → string
	// %d → integer
	// %f → float
	// %v → default representation (works for almost anything)
	var userName = "Rwind"
	var email = "Rwind@gmail.com"

	// fmt.Println("Hello", userName, "and you email is", email, ", Welcome to my project!!!")
	fmt.Printf("Hello %v and you email is %v, Welcome to my project\n", userName, email)


	// String Initialization
	// Go is statically typed language for better storage efficiency and improve perfomance more compared other languages
	
	var movieName string
	var noOfTickets int

	movieName = "Avatar"
	noOfTickets = 3
	
	// Then how this worked var userName = "Rwind" ?? -> Because Go uses compile-time type inference, the compiler infers "Rwind" as a string, so userName is statically typed as string.

	fmt.Printf("you have booked %d for %s movie\n", noOfTickets, movieName)

	// Sytantical Sugar in Go
	// We can also describe a variable like
	// := declares and initializes a new variable (with type inferred) in one step, while "=" only assigns a value to an already-declared variable.

	password := "Rwind@21"
	fmt.Printf("Your password %v", password)
}