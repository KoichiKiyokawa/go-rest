package main

import (
	"fmt"
	"log"
  "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	fmt.Fprintf(w, "<h1>%s</h1>", title)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
