package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hangman", Handler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET")
	case "POST": // Gestion d'erreur
		if err := r.ParseForm(); err != nil {
			return
		}
	}
	variable := r.Form.Get("input")
	fmt.Println(variable)
}
