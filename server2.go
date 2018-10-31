package main

import (
	"net/http"
	"log"
	"fmt"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "url.path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()

	fmt.Fprintf(w, "url.path = %q\n", r.URL.Path)
	mu.Unlock()
}
