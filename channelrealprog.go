package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TO_FU_PRICE float32 = 5

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second + 1)
		var checkPrice = rand.Float32() * 20
		if checkPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second + 1)
		var checkPrice = rand.Float32() * 20
		if checkPrice <= MAX_TO_FU_PRICE {
			tofuChannel <- website
		}
	}
}

// Execute one of the select function and stop
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// This select statement will execute one of the following channels
	select {
	case website := <-chickenChannel:
		fmt.Printf("\n Found a deal on chicken at %v ", website)
	case website := <-tofuChannel:
		fmt.Printf("\n Found a deal on tofu at %v ", website)
	}

}
