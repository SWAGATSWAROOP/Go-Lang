package main

import (
	"errors"
	"fmt"
)

func main() {
	myString := "Swagat"
	printMe(myString)
	var value, err = divide(11, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(value)
	}
	// We can also use the switch block for the above
	switch { //We do not need to specify break in Go
	case err != nil:
		fmt.Println(err.Error())

	case value != 0:
		fmt.Println(value)

	default:
		fmt.Println("No error")
	}
	switch value {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}
	var result, remainder int32 = multipleAndRemainder(10, 3)
	fmt.Println(result, remainder)
	fmt.Printf("The result of multiplication %v and remainder %v \n", result, remainder)
}

func printMe(print string) {
	fmt.Println(print)
}

func divide(a int32, b int32) (int32, error) {
	var err error //Default value of error is nil
	if b == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, err
	}
	return a / b, err
}

// Return multiple values
func multipleAndRemainder(num1 int32, num2 int32) (int32, int32) {
	return num1 * num2, num1 % num2
}
