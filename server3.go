package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "header[%q] = %q \n ", k, v)

	}
	fmt.Fprintf(w, "Host = %q \n ", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q \n ", r.RemoteAddr)
	if err := r.ParseForm(); r != nil {
		log.Print(err)
	}

	for key, value := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", key, value)
	}
}
