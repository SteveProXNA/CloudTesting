package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "5000"
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	http.Handle("/", r)
	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
