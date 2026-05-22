package main

import "fmt"

func main(){
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