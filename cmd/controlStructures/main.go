package main
import "fmt"

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
	calc(2,5,"+")
}