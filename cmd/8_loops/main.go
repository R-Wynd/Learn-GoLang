package main

import "fmt"

func main(){
	// Traditional for loop

	// for i := 0; i<=10; i++{
	// 	fmt.Println(i)
	// }

	// Go dont have while loops, but that can be achieved by for keyword itself
	// i:= 0
	// for{
	// 	if i<= 10{
	// 		fmt.Println(i)
	// 	}else{
	// 		break
	// 	}
	// 	i+=1
	// }

	// Special Go loop:=  Range Loops
	// To iterate over collections like arrays, slices, maps, or strings, use the range keyword. It returns both the index and the value

	cars := []string{"ford", "bmw", "ferrari", "lamborgini"}

	for index, car := range cars{
		fmt.Printf("The index of %v is %v\n", car, index)
	}


	grades := map[string]string{
		"Aravind" : "A",
		"Nitthi" : "S",
		"Rahul" : "c",
	}

	for name, grade := range grades{
		fmt.Printf("%s secured %s grade\n", name, grade)
	}
	

	gradesWithSubs := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"Science": 88,
		},
		"Bob": {
			"Math":    72,
			"History": 90,
		},
	}

	for name, subs := range gradesWithSubs{
		for sub , mark := range subs{
			fmt.Printf("%v scored %v marks in %v\n", name, mark, sub)
		}
	}
}
