package hangmanweb

import (
	"encoding/json"
	hc "hangmanweb/hangman-classic"
	"log"
	"os"
)

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
		SaveData(data.Username, data.Mot, data.Tries, data.Attempts)
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

func LetterPresentStr(word, letter string) bool {
	for i := 0; i < len(word); i++ {
		if string(word[i]) == letter {
			return true
		}
	}
	return false
}

func SaveData(username string, mot string, tries int, attempts int) {
	filename := "./scoreboard/score.txt"
	Data := DataHangman{Username: username, Mot: mot, Tries: tries, Attempts: attempts}
	data, err := json.Marshal(Data)
	if err != nil {
		log.Fatal(err)
	} else {
		file, err := os.OpenFile(filename, os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.Write(data)
		file.Write([]byte("\n"))
	}
}
