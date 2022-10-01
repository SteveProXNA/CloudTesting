package main

import (
	"github.com/nikandfor/goid"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		time.Sleep(1 * time.Second)
		recd := <-ch
		log.Printf("[%d] goroutine recd '%d'", goid.ID(), recd)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		time.Sleep(2 * time.Second)
		send := 42
		ch <- send
		log.Printf("[%d] goroutine send '%d'", goid.ID(), send)
		wg.Done()
	}(ch, wg)

	log.Printf("[%d] running.", goid.ID())
	wg.Wait()
	log.Printf("[%d] COMPLETE", goid.ID())
}
