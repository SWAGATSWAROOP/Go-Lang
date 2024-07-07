package main

import (
	"fmt"
)

type gasEngine struct {
	mpg          uint8
	gallons      uint8
	ownernerInfo owner
}

type owner struct {
	name string
}

// Now the struct has this internal 
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownernerInfo: owner{"Alex"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownernerInfo.name)

	// This struct is not reusable
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myEngine2)

	fmt.Println("Total miles left in tank : %v",myEngine.milesLeft())
}
