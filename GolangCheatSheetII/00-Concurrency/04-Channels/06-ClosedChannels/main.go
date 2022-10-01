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
		// receiving from the channel
		for msg := range ch {
			log.Printf("[%d] goroutine recd '%d'", goid.ID(), msg)
		}
		wg.Done()
	}(ch, wg)

	// Send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		time.Sleep(2 * time.Second)
		for i := 0; i < 3; i++ {
			// sending to the channel
			ch <- i
			log.Printf("[%d] goroutine send '%d'", goid.ID(), i)
		}
		close(ch)
		wg.Done()
	}(ch, wg)

	log.Printf("[%d] running.", goid.ID())
	wg.Wait()
	log.Printf("[%d] COMPLETE", goid.ID())
}
