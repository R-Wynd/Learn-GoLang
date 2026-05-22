package main
import (
	"fmt"
	"errors"
)


func intDiv(numerator int, denominator int) (int, int, error){
	// return a/b, a%b

	// Exception handling
	var err error //error is Go’s built-in type for “something went wrong.”
	if denominator == 0{
		err = errors.New("Cannot divide by zero")
		return 0,0,err
	}

	quotient := numerator/denominator
	reminder := numerator%denominator
	return quotient, reminder, err
}


func calc(a int,b int, op string){
	switch op {
		case "+":
			fmt.Println(a+b)
		case "-":
			fmt.Println(a-b)
		case "*":
			fmt.Println(a*b)
		case "/":
			fmt.Println(a/b)
		default:
			fmt.Println("Invalid operation")
	}
}



func main(){
	// calc(2,5,"+")
	
	a:= 983987345
	b := 2
	quotient, reminder, err := intDiv(a,b)

	if err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Printf("Quotient: %v and Reminder: %v for numerator: %v and denominator: %v\n", quotient, reminder, a,b)
	}

	
}

