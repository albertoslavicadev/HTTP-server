package main

import (
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  string `json:"pages"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content.type", "text/html")

	w.Write([]byte("<h1 style='color: steelblue'>Hello!!!!</h1>"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	http.HandleFunc("/hello", Hello)

	// func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Write([]byte("Hello"))
	// }

	log.Fatal(http.ListenAndServe(":5100", nil))
}
