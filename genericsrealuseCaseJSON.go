package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

// Generics in struct
type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type gasEngine struct {
	gallons float32
	mpkwh   float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("./json/contactinfo.json")
	fmt.Println(contacts)
	fmt.Printf("\n \n")
	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./json/purchaseInfo.json")
	// fmt.Println(purchases)

	for _, i := range purchases {
		fmt.Printf("Name: %v, Price: %v, Amount: %v \n", i.Name, i.Price, i.Amount)
	}

	var gascar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpkwh:   40,
		},
	}
	fmt.Println(gascar)

}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	var data, _ = os.ReadFile(filePath)
	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}
