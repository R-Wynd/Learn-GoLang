package main

import "fmt"

func main(){

	// Arrays in Go
	// An array is a fixed-size collection of elements of the same type which is indexable.
	// Values not filled in the array will print 0 for int and space for string type.
	// 

	// Array Declaration
	var arr = [50]int{1,2,3,4,5}
	var arrStr = [50]string{"a", "b", "c"}

	arr[49] = 100

	fmt.Println(arr)
	fmt.Println(arrStr)

	// Array Initialization
	// var arr2 [10]int     //Method 1
	// var arr2 = [10]int{} //Method 2
	arr2 := [10]int{}       //Method 3
	arr2[0] = 100
	arr2[1] = 200
	fmt.Println(arr2)

	arr3 := [...]int{1,2,3,4} //Method 4, Automatically size infered like have 4 elements
	// arr3[100] = 1 // index 100 out of bounds [0:4]; Couldn't add elements cuz fixed 
	fmt.Println(arr3)


	// A slice in Go is a flexible, dynamically-sized view into the elements of an underlying array. 
	// Unlike arrays, which have a fixed size, slices can grow and shrink as needed, making them the most common way to handle sequences of data in Go


	// var s []int
	s := []int{1,2}
	// fmt.Printf("The len of s is %v and capacity of s is %v\n", len(s), cap(s))
	s = append(s, 100,200,300)
	fmt.Println(s)
	

	// When appending elements, The slice add extra rooms to grows which is called capacity and total size is len. 
	fmt.Printf("The len of s is %v and capacity of s is %v\n", len(s), cap(s))

	// We could't access the index of not assigned capacity 
	// fmt.Println(s[6]) //runtime error: index out of range [6] with length 5



	fmt.Println(s[2:]) // Slice Slicing
	fmt.Println(len(s))
}
