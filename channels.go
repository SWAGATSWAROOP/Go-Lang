package main

import (
	"fmt"
	"time"
)

// Channels : 1.Hold data 2.Thread Safe 3.Listen For Data

// func main() {
// 	var c = make(chan int)
// 	c <- 1 //writing in an unbuffered channel we will be waiting here until somebody reads from it
// 	// Unable to reach the line from below where we read from the channel
// 	// Luckily runtime is smart enough to get this error
// 	var i = <-c
// 	fmt.Println(i)
// }

// Go routines + channel
// func main(){
// 	var c = make(chan int)
// 	go process(c)
// 	fmt.Println(<-c) //reading from the channel
// }

// func process(c chan int){
// 	c <- 123
// }

func main() {
	// Whenever next time visit see removing 5 and adding ot also
	var c = make(chan int, 5) //Adding a buffer of 5 now it will exit early
	go process(c)
	for i := range c { //we will get deadlock error as after channel data is finished it will wait on top
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func process(c chan int) {
	defer close(c) //This means use this before the function exits
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exited")
	// close(c) //This notifies that channel has been closed and we will not get deadlock error
}
