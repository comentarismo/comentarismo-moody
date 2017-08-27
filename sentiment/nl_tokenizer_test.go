package sentiment_test

import (
	"comentarismo-moody/sentiment"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNLTokenize(t *testing.T) {
	tokenizer := sentiment.NewNLTokenizer()
	text := "Super!!!!!! Super handig, ideaal.........nu kunnen we heerlijk spotify luisteren via onze versterker! Echt heel blij met dit handige kastje!"
	m := tokenizer.Tokenize(text)
	log.Printf("%v", m)

	testWords := map[string]int64{
		"handig":  1,
		"lijk":    1,
		"spotify": 1,
		"echt":    1,
		"blij":    1,
		"handi":   1,
		"kast":    1,
	}

	//log.Println(m)

	for testWord, testWordValue := range testWords {
		valid := false
		for foundWord, foundWordValue := range m {
			if testWord == foundWord {
				//log.Println(testWord, testWordValue, foundWord, foundWordValue)
				assert.Equal(t, testWordValue, foundWordValue)
				valid = true
			}
		}
		if !valid {
			t.Fatal("Word expected never found --> ", testWord)
		}
	}

}
