package foo

import (
	"github.com/nikandfor/goid"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	trackChan   = make(chan string)
	untrackChan = make(chan string)
	DoneFunc    = func() {}
)

func Track(hs ...string) {
	for _, h := range hs {
		log.Printf("[%d] Track '%v'", goid.ID(), h)
		trackChan <- h
	}
}

func Untrack(hs ...string) {
	for _, h := range hs {
		log.Printf("[%d] Untrack '%v'", goid.ID(), h)
		untrackChan <- h
	}
}

var hashtags []string

func init() {
	go func() {
		for {
			select {
			case h := <-trackChan:
				log.Printf("[%d] select Track '%v'", goid.ID(), h)
				if !sliceContains(hashtags, h) {
					log.Printf("[%d] select hashtags append '%v'", goid.ID(), h)
					hashtags = append(hashtags, h)
					log.Printf("[%d] select hashtags delays", goid.ID())
					time.Sleep(2 * time.Second)
					DoneFunc()
				}
			case h := <-untrackChan:
				log.Printf("[%d] select Untrack '%v'", goid.ID(), h)
				if sliceContains(hashtags, h) {
					hashtags = sliceRemove(hashtags, h)
					DoneFunc()
				}
			}
		}
	}()
}

// https://stackoverflow.com/questions/34070369/removing-a-string-from-a-slice-in-go
func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// https://stackoverflow.com/questions/34070369/removing-a-string-from-a-slice-in-go
func sliceRemove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
