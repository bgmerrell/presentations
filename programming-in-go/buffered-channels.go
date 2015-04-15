package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

const MAX_CONCURRENT = 3

var ch chan bool = make(chan bool, MAX_CONCURRENT)

var info struct {
	sync.Mutex
	count int
}

func handle() {
	ch <- true

	info.Lock()
	info.count += 1
	fmt.Println(info.count)
	info.Unlock()

	time.Sleep(time.Duration(rand.Intn(500) + 100) * time.Millisecond)

	info.Lock()
	info.count -= 1
	fmt.Println(info.count)
	info.Unlock()

	<-ch
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 20; i++ {
		go handle()
	}
}
