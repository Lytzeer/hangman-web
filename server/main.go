package main

import (
	"fmt"
	hw "hangmanweb"
	hc "hangmanweb/hangman-classic"
	"html/template"
	"net/http"
)

var data hw.Hangman

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", HandleIndex)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/user", HandlerUser)
	http.HandleFunc("/win", HandlerWin)
	http.HandleFunc("/loose", HandlerLoose)
	http.HandleFunc("/hangman", Handler)
	http.HandleFunc("/reset", HandlerReset)
	http.ListenAndServe(":8080", nil)
	return
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if data.Username == "" {
		data.MotTab, data.Mot, data.Motstr = hw.Initword()
		data.Attempts = 10
		http.Redirect(w, r, "/user", 302)
	}
	var tmpl *template.Template
	if data.Win {
		http.Redirect(w, r, "/win", 302)
		return
	} else if data.Attempts <= 0 {
		http.Redirect(w, r, "/loose", 302)
		return
	} else {
		tmpl = template.Must(template.ParseFiles("./static/play.html"))
	}
	tmpl.Execute(w, data)
	return
}

func HandlerUser(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/user.html"))
	tmpl.Execute(w, nil)
	return
}

func HandlerWin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/win.html"))
	if r.Method == "POST" {
		data.MotTab, data.Mot, data.Motstr = hw.Initword()
		data.Attempts = 10
		data.Tries = 0
		data.Win = false
		data.LettersUsed = []string{}
		data.LettersUsedStr = ""
		http.Redirect(w, r, "/", 302)
	}
	tmpl.Execute(w, nil)
	return
}

func HandlerLoose(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/loose.html"))
	if r.Method == "POST" {
		data.MotTab, data.Mot, data.Motstr = hw.Initword()
		data.Attempts = 10
		data.Tries = 0
		data.Win = false
		data.LettersUsed = []string{}
		data.LettersUsedStr = ""
		http.Redirect(w, r, "/", 302)
	}
	tmpl.Execute(w, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	variable := r.FormValue("input")
	if len(variable) == 1 {
		data.LettersUsedStr += variable
	}
	if data.Username == "" {
		data.Username = variable
	} else {
		data.Letter = variable
		data.Tries++
	}
	HangmanWeb()
	http.Redirect(w, r, "/", 302)
}

func HangmanWeb() {
	Motstr, State := hc.IsInputOk(data.Letter, data.Mot, data.Motstr, &data.LettersUsed)
	data.Motstr = Motstr
	if State == "wordwrong" || State == "wordinvalid" {
		data.Attempts -= 2
	}
	if !(LetterPresent(data.Letter)) {
		data.LettersUsed = append(data.LettersUsed, data.Letter)
	}
	if State == "fail" {
		data.Attempts--
	}
	if data.Mot == data.Motstr || State == "wordgood" {
		data.Win = true
		hw.SaveData(data.Username, data.Mot, data.Tries, data.Attempts)
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

func HandlerReset(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/play.html"))
	data.MotTab, data.Mot, data.Motstr = hw.Initword()
	data.Attempts = 10
	data.Tries = 0
	data.Hang = ""
	data.Win = false
	data.LettersUsed = []string{}
	data.LettersUsedStr = ""
	http.Redirect(w, r, "/", 302)
	tmpl.Execute(w, data)
	return
}
