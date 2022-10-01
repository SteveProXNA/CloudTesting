package main

import (
	"concurrencytesting/book"
	"github.com/nikandfor/goid"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]book.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	cacheCh := make(chan book.Book)
	dbCh := make(chan book.Book)

	log.Printf("[%d] starting.", goid.ID())
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)

		// Send only channel	Query
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- book.Book) {
			if b, ok := queryCache(id, m); ok {
				//log.Printf("[%d] goroutine send Cache '%s'", goid.ID(), b.Title)
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)

		// Send only channel	Database
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- book.Book) {
			if b, ok := queryDatabase(id, m); ok {
				//log.Printf("[%d] goroutine send dBase '%s'", goid.ID(), b.Title)
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		// Receive only channel
		go func(cacheCh, dbCh <-chan book.Book) {

			select {
			case b := <-cacheCh:
				log.Printf("[%d] goroutine recd Cache '%s'", goid.ID(), b.Title)
				<-dbCh

			case b := <-dbCh:
				log.Printf("[%d] goroutine recd dBase '%s'", goid.ID(), b.Title)
			}

		}(cacheCh, dbCh)

		time.Sleep(150 * time.Millisecond)
	}

	log.Printf("[%d] running.", goid.ID())
	wg.Wait()
	log.Printf("[%d] COMPLETE", goid.ID())
}

func queryCache(id int, m *sync.RWMutex) (book.Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (book.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			return b, true
		}
	}

	return book.Book{}, false
}
