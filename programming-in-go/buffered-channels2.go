package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

const maxConcurrent = 3

// START OMIT
var ch chan bool = make(chan bool, maxConcurrent)
var wg sync.WaitGroup

var info struct {
	sync.Mutex
	count int
}

func handle() {
	ch <- true
	info.Lock(); info.count += 1; fmt.Println(info.count); info.Unlock()
	time.Sleep(time.Duration(rand.Intn(500) + 100) * time.Millisecond)
	info.Lock(); info.count -= 1; fmt.Println(info.count); info.Unlock()
	<-ch; wg.Done()
}

func main() {
	rand.Seed(42)
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go handle()
	}
	wg.Wait()
}
// END OMIT
