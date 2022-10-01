package main

import (
	"concurrencytesting/book"
	"github.com/nikandfor/goid"
	"math/rand"
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

	for i := 0; i < 2; i++ {
		id := rnd.Intn(10) + 1
		log.Printf("[%d] ID:%d\n", goid.ID(), id)

		go func(id int) {
			if b, ok := queryCache(id); ok {
				log.Printf("[%d] from cache\n", goid.ID())
				log.Printf("[%d] book.Book '%v'\n", goid.ID(), b.Title)
			}
		}(id)

		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				log.Printf("[%d] from database\n", goid.ID())
				log.Printf("[%d] book.Book '%v'\n", goid.ID(), b.Title)
			}
		}(id)
	}

	time.Sleep(2 * time.Second)
}
