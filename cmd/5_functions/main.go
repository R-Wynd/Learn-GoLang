package main

import (
	"fmt"
)

// Basic syntax of functions in Go
// func functionName(param1 type, param2 type) returnType {
//     // body
//     return value
// }


func intDiv(a int, b int) (int, int){
	// return a/b, a%b
	quotient := a/b
	reminder := a%b
	return quotient, reminder
}

func main(){
	a:= 983987345
	b := 657
	quotient, reminder := intDiv(a,b)
	fmt.Printf("Quotient: %v and Reminder: %v for numerator: %v and denominator: %v\n", quotient, reminder, a,b)
}