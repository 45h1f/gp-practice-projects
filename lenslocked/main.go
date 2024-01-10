package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, World ok!</h1>")
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About page!</h1>")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact page!</h1>")
}

func faqPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>FAQ page!</h1>")
}
func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/faq", contactPage)
	fmt.Println("Starting server on the port 8080...")
	http.ListenAndServe(":8080", nil)
}
