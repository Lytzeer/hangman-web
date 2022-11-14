package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Hangman struct {
	Letter    string
	NbLetters []string
}

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", Handler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/hangman", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST": // Gestion d'erreur
		if err := r.ParseForm(); err != nil {
			return
		}
	}

	// Récupérez votre valeur
	variable := r.Form.Get("input")
	//fmt.Println(variable)
	tab := []string{}
	tab = append(tab, "let")
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	data := Hangman{
		Letter:    variable,
		NbLetters: tab,
	}
	fmt.Println(data)
	tmpl.Execute(w, data)

}
