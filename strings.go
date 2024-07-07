package main

import "fmt"

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\n The length of myString is %v", len(myString))
	var myStr = []rune("résumé")
	var index = myStr[1]
	fmt.Printf("%v %T\n", index, index)
	for i, v := range myStr {
		fmt.Println(i, v)
	}

	var myRune = 'a'
	fmt.Printf("\n myRune = %v", myRune)

	

	// Strings are immutable in Go
	// go deals with string as bytes and not number of characters
	//rune are unicode point numbers that point to characters
}
