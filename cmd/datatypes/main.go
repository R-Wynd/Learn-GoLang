package main

import "fmt"

func main(){
	// Basic Primitive types 
	// Category	Types
	// Boolean	bool
	// Signed Integer	int, int8, int16, int32, int64
	// Unsigned Integer	uint, uint8, uint16, uint32, uint64, uintptr
	// Floating-Point	float32, float64
	// Complex	complex64, complex128
	// Text	string, rune (int32 alias), byte (uint8 alias)

	// max of int32 is 32,767, So using "int" alone, complier will chose 32 or 64 bits based on the value.
	var intNum int = 32767
	intNum += 1
	fmt.Println(intNum)

	// Like int we cant use float directly, we need to mention what what float we gonna use like float32 or float64
	var floatNum32 float32 = 12345678.9
	fmt.Println(floatNum32) // 1.2345679e+07; Println chooses scientific notation when the number is big
	fmt.Printf("%f\n", floatNum32) // 12345679.000000

	var floatNum64 float64 = 12345678.9
	fmt.Printf("%f\n", floatNum64) //12345678.900000

	//float32 has ~7 decimal digits of precision; 12345678.9 needs more, so it rounds to 12345679. 
	//float64 has enough precision, so you get 12345678.900000




	// Can't do arth op with two diff data types, For that we need to type cast any one data type 
	// var result = floatNum64 + intNum //invalid operation: floatNum64 + intNum (mismatched types float64 and int)

	result := floatNum64 + float64(intNum)
	fmt.Printf("Addition: %v\n", result)


	myStr := "Hello world"

	fmt.Println(len(myStr))
}