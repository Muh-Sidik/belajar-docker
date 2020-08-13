package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var hello = []byte("Hello")
		fmt.Fprintf(w, string(hello))
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
