package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))


}

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"url.path = %q\n",r.URL.Path)
}
