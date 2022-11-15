package hangmanweb

import (
	hc "hangmanweb/hangman-classic"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func Initword(filename string) ([]string, string, string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	var hold string
	list := []string{}
	for _, m := range string(data) {
		if m != 10 {
			hold = hold + string(m)
		} else {
			if hold != "" {
				list = append(list, hold)
				hold = ""
			}
		}
	}

	rand.Seed(time.Now().UnixNano())

	randomword := list[rand.Intn(len(list))]

	randomwordhide := hc.CreateWord(randomword)
	randomwordhide = randomwordhide[:(len(randomwordhide) - 1)]
	randomwordhidetab := []string{}
	for i := 0; i < len(randomwordhide)-1; i++ {
		randomwordhidetab = append(randomwordhidetab, string(randomwordhide[i]))
	}

	randomword = randomword[:(len(randomword) - 1)]

	return randomwordhidetab, randomword, randomwordhide

}
func TabtoStr(tab []string) string {
	str := ""
	for _, ch := range tab {
		str += string(ch)
	}
	return str
}
