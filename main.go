package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hangman", Handler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
}
