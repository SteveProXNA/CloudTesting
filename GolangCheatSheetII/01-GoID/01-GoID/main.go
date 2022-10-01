package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nikandfor/goid"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		_, _ = fmt.Fprintln(os.Stdout, goid.ID(), s)
	}
}

func main() {

	go say("world")
	say("hello")
}
