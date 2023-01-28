package main

import (
	"fmt"
	hw "hangmanweb"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", hw.HandleIndex)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/user", hw.HandlerUser)
	http.HandleFunc("/win", hw.HandlerWin)
	http.HandleFunc("/loose", hw.HandlerLoose)
	http.HandleFunc("/hangman", hw.Handler)
	http.HandleFunc("/reset", hw.HandlerReset)
	http.ListenAndServe(":8080", nil)
	return
}
