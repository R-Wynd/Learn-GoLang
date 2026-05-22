package main

import "fmt"

func main(){
	// Pointer is a variable that stores the memory address of another value, rather than the value itself

	// Why Use Pointers?

	// Modify Data Directly: By default, Go passes arguments to functions by value (it makes a copy). If you want a function to change the original variable, you must pass a pointer to it.
	// Performance (Efficiency): For very large data structures (like huge structs), passing a copy can be slow and memory-intensive. Passing a pointer is always small (usually 8 bytes on a 64-bit system).
	// Pointer Receivers: Methods in Go often use pointer receivers so they can modify the original fields of the struct they are attached to
	// Unlike languages like C, Go does not allow you to add or subtract from pointers (e.g., ptr++ is not allowed). This makes Go safer and less prone to memory corruption.
	// An uninitialized pointer defaults to nil. Attempting to dereference a nil pointer will cause your program to panic (crash).


	age := 25
	var ptx *int = &age
	

	fmt.Println(ptx)   //Memory address of age: 0x14000010090
	fmt.Println(*ptx)  //Orginal value of age 25
	fmt.Println(&ptx)  //Memory address of ptx stored: 0x14000052038


	// Modify Data Directly: By default, Go passes arguments to functions by value (it makes a copy). If you want a function to change the original variable, you must pass a pointer to it.

	num := 4

    squareCopy(num)
    fmt.Println(num) // Prints: 4 (unchanged)

    squarePointer(&num)
    fmt.Println(num) // Prints: 16 (changed)



	var newptx *int
	fmt.Println(newptx) //Safe <nil> printed
	// fmt.Println(*newptx) //Crash -> invalid memory address or nil pointer dereference
	fmt.Println(&newptx) // Memory address of newptx stored


	slice := []int{1,2,3}
	slicecopy := slice
	slicecopy[2] = 4
	fmt.Println(slice)     //[1 2 4]
	fmt.Println(slicecopy) //[1 2 4]

	// The original slice also changed expected slicecopy alone need to be changed. But understand the hood slice uses pointers. So the original values changed when reusing. 


}


func squareCopy(n int) {
    n = n * n // Modifies a copy, original stays safe
}

func squarePointer(n *int) {
    *n = *n * *n // Modifies the original value directly
}