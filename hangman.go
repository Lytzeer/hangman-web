package hangmanweb

import hc "hangmanweb/hangman-classic"

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
		//SaveData(data.Username, data.Mot, data.Tries, data.Attempts)
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
