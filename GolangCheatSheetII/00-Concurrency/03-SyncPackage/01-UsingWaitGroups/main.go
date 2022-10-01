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

func queryCache(id int) (book.Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (book.Book, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}

	return book.Book{}, false
}

func main() {

	wg := &sync.WaitGroup{}
	for i :=0; i < 2; i++ {
		id := rnd.Intn(10) + 1
		log.Printf("[%d] ID:%d\n", goid.ID(), id)

		wg.Add(2)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryCache(id); ok {
				log.Printf("[%d] from cache\n", goid.ID())
				log.Printf("[%d] Book '%v'\n", goid.ID(), b.Title)
			}
			wg.Done()
		}(id, wg)


		go func(id int, wg *sync.WaitGroup) {
			if b, ok :=queryDatabase(id); ok {
				log.Printf("[%d] from database\n", goid.ID())
				time.Sleep(2 * time.Second)
				log.Printf("[%d] Book '%v'\n", goid.ID(), b.Title)
				time.Sleep(2 * time.Second)
			}
			wg.Done()
		}(id, wg)
	}

	//time.Sleep(2 * time.Second)
	log.Printf("[%d] running.\n", goid.ID())
	wg.Wait()
	log.Printf("[%d] COMPLETE\n", goid.ID())
}