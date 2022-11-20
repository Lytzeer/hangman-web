package hangmanweb

import (
	"html/template"
	"net/http"
)

var data Hangman

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if data.Username == "" {
		data.MotTab, data.Mot, data.Motstr = Initword()
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
		data.MotTab, data.Mot, data.Motstr = Initword()
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
		data.MotTab, data.Mot, data.Motstr = Initword()
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
		if LetterPresentStr(data.LettersUsedStr, variable) == false {
			data.LettersUsedStr += variable
		}
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

func HandlerReset(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/play.html"))
	data.MotTab, data.Mot, data.Motstr = Initword()
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
