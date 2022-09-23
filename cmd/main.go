package main

import (
	"fmt"
	"log"
	"time"

	"github.com/matti/fan"
)

func subscriber(amount int, f *fan.Fan) {
	ch := make(chan string)
	f.Subscribe <- ch

	for i := 0; i < amount; i++ {
		msg := <-ch
		log.Println("receive", "got", msg)
		time.Sleep(time.Second * 2)
	}
	log.Println("unsub", amount)
	f.Unsubscribe <- ch
}

func main() {

	ch := make(chan string)

	f := fan.Run(ch)

	go subscriber(3, f)
	go subscriber(9, f)

	for i := 0; ; i++ {
		f.Publish <- fmt.Sprintf("%d", i)
		time.Sleep(time.Second)
	}
}
