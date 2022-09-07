package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	bind := ":8081"
	log.Println("Starting web server on port", bind)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Hi, %q", html.EscapeString(r.URL.Path))
		log.Println(msg)
		_, err := fmt.Fprint(w, msg)
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	log.Fatal(http.ListenAndServe(bind, nil))
}
