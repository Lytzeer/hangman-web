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
	Mot         string
	Attempts    int
	LettersUsed []string
}

var data Hangman

func main() {
	data.MotTab, data.Mot = hw.Initword(os.Args[len(os.Args)-1])
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
	data.LettersUsed = append(data.LettersUsed, variable)
	fmt.Println(data)
	tmpl.Execute(w, data)
	hc.IsInputOk(data.Letter, hw.TabtoStr(data.Mot))
}

func SetData(data *Hangman, letter string) {
	data.Letter = letter
}
