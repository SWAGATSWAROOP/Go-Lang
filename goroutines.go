package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// We should use mutex for concurrency control on same meemory
var m = sync.RWMutex{} //Mutual Exclusion

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}
var wg = sync.WaitGroup{} //Wait group for waiting for response from go routines
func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// Counters like WaitGroup
		wg.Add(1) //Increments the counter by 1
		// for parallel calls use go keyword go for go routines
		go dbCall(i)
	}
	wg.Wait() //Now we will wait
	fmt.Printf("\n Total execution time: %v \n", time.Since(t0))
	fmt.Println(results)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock() //Locks if already being used
	results = append(results, dbData[i])
	m.Unlock() //unlocks after critical section has been performed
	save(dbData[i])
	log()
	wg.Done()  //Decrements the counter
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Printf("\n the current results are: %v",results)
	m.RUnlock()
}
