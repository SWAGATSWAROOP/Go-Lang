package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	var floatSlice = []float32{1.1, 2.2, 3.4}
	fmt.Println(sumSlice[float32](floatSlice))

	fmt.Println(checkLengthIsZero(intSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func checkLengthIsZero[T any](slice []T) bool {
	return len(slice) == 0
}
