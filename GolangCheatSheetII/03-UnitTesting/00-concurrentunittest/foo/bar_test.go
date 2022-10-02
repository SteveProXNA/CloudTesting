package foo

import (
	"github.com/nikandfor/goid"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestTrack(t *testing.T) {

	log.Printf("[%d] TestTrack", goid.ID())
	doneChan := make(chan struct{}, 1)
	DoneFunc = func() {
		log.Printf("[%d] DoneFunc doneChan send", goid.ID())
		doneChan <- struct{}{}
	}
	defer func() {
		log.Printf("[%d] DoneFunc defer", goid.ID())
		DoneFunc = func() {}
	}()

	h := "name"

	Track(h)
	log.Printf("[%d] doneChan wait", goid.ID())
	<-doneChan
	log.Printf("[%d] doneChan recd", goid.ID())

	a := sliceContains(hashtags, h)
	if !a {
		t.Errorf("not contains")
	}
}

func TestUntrack(t *testing.T) {
	doneChan := make(chan struct{}, 1)
	DoneFunc = func() {
		doneChan <- struct{}{}
	}
	defer func() {
		DoneFunc = func() {}
	}()

	name := "name"
	hashtags = []string{name, "surname"}

	Untrack(name)
	<-doneChan

	a := sliceContains(hashtags, name)
	if a {
		t.Errorf("not contains")
	}
}

func TestTrackUntrack(t *testing.T) {

	doneChan := make(chan struct{}, 1)
	DoneFunc = func() {
		doneChan <- struct{}{}
	}
	defer func() {
		DoneFunc = func() {}
	}()

	h := "name"

	Track(h)
	<-doneChan

	Untrack(h)
	<-doneChan

	a := sliceContains(hashtags, h)
	if a {
		t.Errorf("not contains2")
	}
}
