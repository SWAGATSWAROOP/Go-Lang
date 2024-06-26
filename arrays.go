package main

import "fmt"

func main() {
	var arr [3]int32 = [3]int32{1, 2, 3}
	// We can also use the short hand declaration
	arr2 := [...]int32{1, 3} //This will automatically detect the length of the array
	fmt.Println(arr2)
	fmt.Println(arr)
	fmt.Println(arr[1:3]) //This is called slicing
	fmt.Println(arr[1:])
	fmt.Println(arr[:2])
	fmt.Println(&arr[0]) //This will give the address of the first element of the array

	// Arrays with variable length using Slices which are wrapper around array
	var slice []int32 = []int32{1, 2, 3}
	fmt.Println(slice)
	fmt.Printf("Length before append %v and capacity before append %v \n", len(slice), cap(slice))
	slice = append(slice, 4)
	fmt.Printf("Length after append %v and capacity before append %v \n", len(slice), cap(slice))

	// We can also add multiple elements using the spread operator
	slice = append(slice, slice...)
	fmt.Println(slice)

	// We can also add multiple elements normally
	slice = append(slice, 5, 6, 7)
	fmt.Println(slice)

	// We can also create a slice using make
	slice2 := make([]int32, 5, 20)    //This will create a slice of length 5 and capacity 8(optional) this helps in performance improve as reallocation is not needed
	slice2 = append(slice2, slice...) //By default the capacity would be equal to the length of the slice if not provided
	fmt.Println(slice2)

	for i, v := range arr {
		fmt.Printf("The index is %v and the value is %v \n", i, v)
	}
}
