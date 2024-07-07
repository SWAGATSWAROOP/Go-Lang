package main

import "fmt"

func square(thing2 *[5]float64) [5]float64 {
	for i := range (*thing2) {
		(*thing2)[i] *= (*thing2)[i];
	}
	return *thing2
}

func main() {
	var p *int32 = new(int32) //Intially null //But if point to a meemory location its automatically set 0 int "" for strings and false for boolean
	fmt.Printf("The vale p points to is : %v \n", *p)
	*p = 10
	fmt.Printf("%v is the address and %v is the value \n", p, *p)
	var arrayFloat = [5]float64{2.4, 5.6, 4, 2, 4}
	fmt.Println(square(&arrayFloat))
}
