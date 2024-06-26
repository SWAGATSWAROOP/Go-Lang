package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int8 = 100 //Signed integer
	var uintNum uint8 = 2 //Unsigned integer
	fmt.Println(intNum, uintNum)

	var floatNum float32 = 1122.45
	var floatNum2 float64 = 2.45
	fmt.Println(floatNum, floatNum2)

	// We cannot add variables of different types even int32 + int64

	var num1 int32 = 100
	var num2 int32 = 3
	fmt.Println(num1+num2, num1%num2)

	// var str string = "Swagat"
	fmt.Println(len("₩")) //But characters outside the ascii character set will be stored as 2 bytes in Go
	// GO uses utf-8 encoding
	fmt.Println(utf8.RuneCountInString("₩")) //This will give the correct length of the string

	var myRune rune = 'W'
	fmt.Println(myRune)

	var myBoolean bool = true

	// Default values for all ints floats and rune is 0 for string it is empty string and for boolean it is false
	fmt.Println(myBoolean)

	// We can also use the short hand declaration
	myString := "Swagat" //Automatically detects the type and := is used for short hand declaration
	fmt.Println(myString)

	var1, var2 := 12, 13
	println(var1, var2)

	// Constants
	const myConst int = 100

}
