package sentiment_test

import (
	"testing"
	"comentarismo-moody/sentiment"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestFrTokenize(t *testing.T) {
	tokenizer := sentiment.NewFrTokenizer()
	text := "Vous avez bien dit en doublant le salaire des ouvriers"
	m := tokenizer.Tokenize(text)

	testWords := map[string]int64{
		"avez":1,
		"dit":1,
		"ire":1,
		"vriers":1,
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
