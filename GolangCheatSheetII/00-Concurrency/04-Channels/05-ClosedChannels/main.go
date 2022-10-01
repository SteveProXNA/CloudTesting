package main

import (
	"github.com/nikandfor/goid"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) // buffered channel

	wg.Add(2)

	// Receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		time.Sleep(1 * time.Second)
		if msg, ok := <-ch; ok {
			log.Printf("[%d] goroutine recd '%d' [%v]", goid.ID(), msg, ok)
		}
		wg.Done()
	}(ch, wg)

	// Send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		time.Sleep(2 * time.Second)
		//ch <- 0
		close(ch)
		wg.Done()
	}(ch, wg)

	log.Printf("[%d] running.", goid.ID())
	wg.Wait()
	log.Printf("[%d] COMPLETE", goid.ID())
}
