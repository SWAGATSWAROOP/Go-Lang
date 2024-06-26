package main

import "fmt"

func main() {
	var myMap map[string]uint32 = map[string]uint32{"one": 1, "two": 2, "three": 3}
	fmt.Println(myMap["one"]) //Accessing the value of the key

	//Also there is boolean variable which returns if the value exists in map of not
	var value, exits = myMap["four"]
	if exits {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exists")
	}

	//We can also delete value from map
	delete(myMap, "one") //This is passed by reference so no return type
	fmt.Println(myMap)

	// For iterating over the map we can use the range keyword which can also be used for arrays and slices
	for key, value := range myMap {
		fmt.Printf("The key is %v and %v is the value \n", key, value) //Order is not preserved in case of maps
	}

}
