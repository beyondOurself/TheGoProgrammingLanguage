package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprint(w, "URL.path = %q\n", r.URL.Path)

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
