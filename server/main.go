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
	http.HandleFunc("/win", HandleIndex)
	http.HandleFunc("/loose", HandleIndex)
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
		tmpl = template.Must(template.ParseFiles("./static/win.html"))
		data.MotTab, data.Mot, data.Motstr = hw.Initword()
		data.Attempts = 10
		data.Win = false
		data.LettersUsed = []string{}
		fmt.Println(data.Mot)
		fmt.Println(data.Attempts)
	} else if data.Attempts == 0 {
		tmpl = template.Must(template.ParseFiles("./static/loose.html"))
		data.MotTab, data.Mot, data.Motstr = hw.Initword()
		data.Attempts = 10
		data.LettersUsed = []string{}
		fmt.Println(data.Mot)
		fmt.Println(data.Attempts)
	} else {
		tmpl = template.Must(template.ParseFiles("./static/play.html"))
	}
	hang := ""
	switch data.Attempts {
	case 1:
		hang = "/static/pics/9.png"
	case 2:
		hang = "/static/pics/8.png"
	case 3:
		hang = "/static/pics/7.png"
	case 4:
		hang = "/static/pics/6.png"
	case 5:
		hang = "/static/pics/5.png"
	case 6:
		hang = "/static/pics/4.png"
	case 7:
		hang = "/static/pics/3.png"
	case 8:
		hang = "/static/pics/2.png"
	case 9:
		hang = "/static/pics/111.png"
	case 10:
		hang = "/static/pics/11.png"
	}
	data.Hang = hang
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
	fmt.Println(data.Mot)
	//fmt.Println(variable)
	// tmpl := template.Must(template.ParseFiles("./static/play.html"))
	if data.Username == "" {
		data.Username = variable
		fmt.Println(data.Username)
	} else {
		data.Letter = variable
		data.Tries++
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
	fmt.Println(data.Mot)
	fmt.Println(data.Attempts)
	http.Redirect(w, r, "/", 302)
	tmpl.Execute(w, data)
	return
}

//Pour la carte d'identité faut rediriger depuis la fonction handler vers /http.HandleFunc("/hangman", Handler)
