package main

import (
	"concurrencytesting/book"
	"github.com/nikandfor/goid"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var cache = map[int]book.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func queryCache(id int, m *sync.Mutex) (book.Book, bool) {
	m.Lock()
	b, ok := cache[id]
	m.Unlock()
	return b, ok
}

func queryDatabase(id int, m *sync.Mutex) (book.Book, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}

	return book.Book{}, false
}

func main() {

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for i :=0; i < 200; i++ {
		id := rnd.Intn(10) + 1
		log.Printf("[%d] ID:%d", goid.ID(), id)

		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			if b, ok := queryCache(id, m); ok {
				log.Printf("[%d] from cache", goid.ID())
				log.Printf("[%d] Book '%v'", goid.ID(), b.Title)
			}
			wg.Done()
		}(id, wg, m)


		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			if b, ok :=queryDatabase(id, m); ok {
				log.Printf("[%d] from database", goid.ID())
				time.Sleep(2 * time.Second)
				log.Printf("[%d] Book '%v'", goid.ID(), b.Title)
				time.Sleep(2 * time.Second)
			}
			wg.Done()
		}(id, wg, m)
	}

	log.Printf("[%d] running.", goid.ID())
	wg.Wait()
	log.Printf("[%d] COMPLETE", goid.ID())
}