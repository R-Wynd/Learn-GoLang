package main 

import "fmt"

func varsLearn(){
	// Variable in Go
	const name = "Aravind"
	var age = 22
	fmt.Println("Hello", name)
	fmt.Println("Your age is", age)


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
	
	// Then how this worked var userName = "Rwind" ?? -> Because Go uses compile-time type inference, the compiler infers "Rwind" is a string, so userName is statically typed as string.

	fmt.Printf("you have booked %d for %s movie\n", noOfTickets, movieName)

	// Sytantical Sugar in Go
	// We can also describe a variable like
	// := declares and initializes a new variable (with type inferred) in one step, while = only assigns a value to an already-declared variable.

	password := "Rwind@21"
	fmt.Printf("Your password %v", password)
}

func userInput(){
	// Getting user Input
	// A pointer in Go is a variable that stores the memory address of another variable, Which is implemented like &varname
	// fmt.Scan(&firstName) passes the memory address of firstName so Scan can store the user’s input directly into that variable.

	var firstName string
	var lastName string
	var movieName string
	var noOfTickets int

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ") 
	fmt.Scan(&lastName)

	fmt.Print("Which movie you would like to watch: ")
	fmt.Scan(&movieName)

	fmt.Print("How many tickets you want: ") 
	fmt.Scan(&noOfTickets)

	fmt.Printf("Hello %s %s, you have booked %d tickets to %s\n", firstName, lastName, noOfTickets, movieName)

	// But we cannot use this for const variables and assinging datatype to it using this syntax
}

func arrLearn(){

	// Arrays in Go
	// An array is a fixed-size collection of elements of the same type.
	// Values not filled in the array will print 0 for int and space for string type

	// Array Declaration
	var arr = [50]int{1,2,3,4,5}
	var arrStr = [50]string{"a", "b", "c"}

	arr[49] = 100

	fmt.Println(arr)
	fmt.Println(arrStr)

	// Array Initialization
	// var arr2 = [10]int{} //Method 1
	// var arr2 [10]int   // Method 2
	arr2 := [10]int{}
	arr2[0] = 100
	arr2[1] = 200
	fmt.Println(arr2)

	arr3 := [...]int{1,2,3,4} // Automatically size infered like have 4 elements
	// arr3[100] = 1 // Could add elements cuz fixed 
	fmt.Println(arr3)
}

func sliceLearn(){
	// A slice is a abstraction of array but its dynamic, flexible will not have any fixed value like array
	// var s []int
	s := []int{1,2}
	s = append(s, 100,200,300)
	fmt.Println(s)
	fmt.Println(s[2:]) // Slice Slicing
	fmt.Println(len(s))
}






func main() {

	// Simple Print statement in Go
	// fmt.Println("Hello World!!")

	// varsLearn()
	// userInput()
	// arrLearn()
	// sliceLearn()
	


}