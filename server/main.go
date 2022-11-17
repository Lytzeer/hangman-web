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
	http.HandleFunc("/", HandleIndex)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/user", HandlerUser)
	http.HandleFunc("/hangman", Handler)
	http.ListenAndServe(":8080", nil)
	return
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if data.Username == "" {
		http.Redirect(w, r, "/user", 302)
	}
	var tmpl *template.Template
	if data.Win {
		tmpl = template.Must(template.ParseFiles("./static/win.html"))
	} else if data.Attempts == 0 {
		tmpl = template.Must(template.ParseFiles("./static/loose.html"))
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
	tmpl.Execute(w, nil)
	return
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// switch r.Method {
	// case "POST": // Gestion d'erreur
	// 	if err := r.ParseForm(); err != nil {
	// 		return
	// 	}
	// }
	// Récupérez votre valeur
	variable := r.FormValue("input")
	//fmt.Println(variable)
	// tmpl := template.Must(template.ParseFiles("./static/play.html"))
	if data.Username == "" {
		data.Username = variable
		fmt.Println(data.Username)
	} else {
		data.Letter = variable
		fmt.Println(data.Username)
	}
	//data.LettersUsed = append(data.LettersUsed, variable)
	fmt.Println(data.Motstr)
	HangmanWeb()
	// tmpl.Execute(w, data)
	http.Redirect(w, r, "/", 302)
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
		data.Win = true
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

//Pour la carte d'identité faut rediriger depuis la fonction handler vers /http.HandleFunc("/hangman", Handler)
