package main

import "fmt"


type gasEngine struct{
	mileage int
	fuel int
	ownerName owner
}

type electricEngine struct{
	mpkwh int
	kwh int
}

type owner struct{
	name string
	drivingLicense string
}

func (ge gasEngine) milesLeft() int{
	return ge.mileage*ge.fuel
}

func (ee electricEngine) milesLeft() int{
	return ee.mpkwh*ee.kwh
}

type engine interface{
	milesLeft() int
}

func canIMakeIt(e engine, distance int) bool{
	if e.milesLeft() < distance{
		return false
	}else{
		return true
	}
}

func main(){

	// A struct (short for structure) in Go is a user-defined composite data type 
	// Groups together variables of different types into a single unit. 
	// It is Go's primary way of modeling real-world entities and is the closest equivalent to "classes" in object-oriented languages.

	var myGasEngine gasEngine = gasEngine{mileage: 40, fuel: 5, ownerName: owner{name: "Alex"}}
	fmt.Println(myGasEngine)

	// Update 
	myGasEngine.mileage = 42
	myGasEngine.fuel = 3
	fmt.Println(myGasEngine)

	fmt.Println(myGasEngine.milesLeft())
	fmt.Println(canIMakeIt(myGasEngine, 100))

	myElectricEngine := electricEngine{mpkwh: 100, kwh: 60}
	fmt.Println(myElectricEngine)
	fmt.Println(canIMakeIt(myElectricEngine, 200))




}
