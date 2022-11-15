package main

import (
	"fmt"
	hw "hangmanweb"
	hc "hangmanweb/hangman-classic"
	"html/template"
	"net/http"
	"os"
)

type Hangman struct {
	Letter      string
	MotTab      []string
	Motstr      string
	Mot         string
	Attempts    int
	LettersUsed []string
}

var data Hangman

func main() {
	data.MotTab, data.Mot, data.Motstr = hw.Initword(os.Args[len(os.Args)-1])
	data.Attempts = 10
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
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	data.Letter = variable
	//data.LettersUsed = append(data.LettersUsed, variable)
	fmt.Println(data.MotTab)
	tmpl.Execute(w, data)
	a, b := hc.IsInputOk(data.Letter, data.Mot, data.Motstr, &data.LettersUsed)
	fmt.Println(a)
	fmt.Println(b)
	return
}
