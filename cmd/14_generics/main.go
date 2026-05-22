package main
import "fmt"

// Can pass any data types
func printAnySlice[T any](slice []T){
	for _, item := range slice{
		fmt.Print(item, " ")
	}
	fmt.Println()
}

// Can pass only int, float32 and float64 data types
func sumSlices[T int | float32 | float64](slice []T)T{
	var sum T
	for _, item := range slice{
		sum += item
	}
	return sum
}

func main(){
	// In Go, Generics allow you to write code (like functions or structs) that can work with multiple data types without sacrificing type safety or performance. 
	// Before Generics were added in Go 1.18, if you wanted a function to work with integers, floats, and strings, you either had to write three identical functions with different names, or use interface{} (which strips away type safety and is slower).


	nums := []int{1,2,3,4,4,5,6}
	strs := []string{"a", "b", "c", "d"}

	printAnySlice(nums)
	printAnySlice(strs)

	fmt.Println(sumSlices(nums))
	// sumSlices(strs) //string does not satisfy int | float32 | float64 (string missing in int | float32 | float64)

	floats := []float64{1.23, 3.2, 3.43, 4.23}
	fmt.Println(sumSlices(floats))


	// generics with structs 
	gasCar := car[gasEngine]{
		carName: "BMW",
		carModel: "M3",
		year: 2018,
		engine : gasEngine{
			mpg: 12,
			gallons: 40,
		},
	}

	fmt.Println(gasCar)


	evCar := car[electricEngine]{
		carName: "BMW",
		carModel: "i3",
		year: 2018,
		engine : electricEngine{
			kwh: 132,
			mpkwh: 23,
		},
	}

	fmt.Println(evCar)
}


// generics with structs 

type gasEngine struct{
	mpg int
	gallons int
}

type electricEngine struct{
	mpkwh int
	kwh int
}

type car[T gasEngine | electricEngine] struct{
	carName string
	carModel string
	year int
	engine T //Need to give engine gasEngine or engine electricEngine. I we want general type we can use generics here
}