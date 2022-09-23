package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/matti/fan"
)

func subscriber(id string, f *fan.Fan, wg *sync.WaitGroup) {
	mySubscription := make(chan string)
	f.Subscribe <- mySubscription
	wg.Done()
	for {
		fmt.Println(id, <-mySubscription)
	}
}

func main() {

	f := fan.Run(
		make(chan string),
	)

	var wg sync.WaitGroup

	wg.Add(1)
	go subscriber("a", f, &wg)
	wg.Add(1)
	go subscriber("b", f, &wg)

	// Wait for subscribers to have subscribed
	wg.Wait()

	for i := 0; ; i++ {
		f.Publish <- fmt.Sprintf("%d", i)
		time.Sleep(time.Second)
	}
}
