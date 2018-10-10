package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/", index)
	fmt.Println("Starting Server ...")
	http.ListenAndServe(":3000", nil)
}
