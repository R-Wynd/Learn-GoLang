package main

import "fmt"

func main(){

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


	// A slice is a abstraction of array but its dynamic, flexible will not have any fixed value like array
	// var s []int
	s := []int{1,2}
	s = append(s, 100,200,300)
	fmt.Println(s)
	fmt.Println(s[2:]) // Slice Slicing
	fmt.Println(len(s))
}
