package main

import "fmt"

func main(){
	// Maps that store unordered collections of key-value pairs. Like a dictionary in python
	
	//Map initialization
	var myMap1 map[string]int = make(map[string]int)   //Method 1
	var myMap2 = make(map[string]int)                  //Method 2
	myMap3 := make(map[string]int)                     //Method 3
	myMap4 := map[string]int{"John": 34, "Alex": 42}   //method 4


	//Adding element in maps 
	myMap1["Aravind"] = 21 
	myMap4["Jack"] = 26

	//Update elements in maps
	myMap1["Aravind"] = 24

	// delete(m, "key")	Removes the entry from the map.
	delete(myMap4, "Jack")
	fmt.Println(myMap4)
	fmt.Println(len(myMap4))


	fmt.Println(myMap1)
	fmt.Println(myMap2)
	fmt.Println(myMap3)
	fmt.Println(myMap4)
	fmt.Println(len(myMap4))

	// The "comma-ok" idiom; ok is true if the key exists, false otherwise.
	val, ok := myMap4["jason"]
	if ok{
		fmt.Println(val)
	}else{
		fmt.Println("The key is not exists.")
	}

	// Nested Maps
	// Map of students with their subject grades
	studentGrades := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"Science": 88,
		},
		"Bob": {
			"Math":    72,
			"History": 90,
		},
	}

	fmt.Println(studentGrades)

}
