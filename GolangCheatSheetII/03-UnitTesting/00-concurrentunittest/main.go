package main

import (
	"fmt"
	"unittesting/foo"
)

func main() {
	fmt.Println("start")

	doneChan := make(chan struct{}, 1)
	foo.DoneFunc = func() {
		doneChan <- struct{}{}
	}
	defer func() {
		foo.DoneFunc = func() {}
	}()

	foo.Track("hello")
	<-doneChan

	foo.Track("world")
	<-doneChan

	foo.Untrack("world")
	<-doneChan

	foo.Untrack("hello")
	<-doneChan

	fmt.Println("finish")
}
