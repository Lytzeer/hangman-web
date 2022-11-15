package main

import (
	"fmt"
	hw "hangmanweb"
	hc "hangmanweb/hangman-classic"
	"html/template"
	"net/http"
	"os"
)

var data hw.Hangman

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
	fmt.Println(data.Motstr)
	HangmanWeb()
	tmpl.Execute(w, data)
	return
}

func HangmanWeb() {
	Motstr, State := hc.IsInputOk(data.Letter, data.Mot, data.Motstr, &data.LettersUsed)
	data.Motstr = Motstr
	if !(LetterPresent(data.Letter)) {
		data.LettersUsed = append(data.LettersUsed, data.Letter)
	}
	if State == "fail" {
		data.Attempts--
	}
	if data.Mot == data.Motstr {
		fmt.Println("gg")
		os.Exit(1)
	}
}

func LetterPresent(letter string) bool {
	for _, ch := range data.LettersUsed {
		if data.Letter == ch {
			return true
		}
	}
	return false
}
