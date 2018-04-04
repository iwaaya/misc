package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	}))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
