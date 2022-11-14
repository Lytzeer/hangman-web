package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hangman", Handler)
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
}
