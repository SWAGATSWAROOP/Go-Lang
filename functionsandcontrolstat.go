package main

import "fmt"

func main() {
	myString := "Swagat"
	printMe(myString)
	println(divide(11, 2))
}

func printMe(print string) {
	fmt.Println(print)
}

func divide(a int32, b int32) int32 {
	return a / b
}
