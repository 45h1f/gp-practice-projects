package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, World!</h1>")
}

func main() {
	http.HandleFunc("/", homePage)

	fmt.Println("Starting server on the port 8080...")
	http.ListenAndServe(":8080", nil)
}
