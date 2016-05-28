package sentiment_test

import (
	"testing"
	"comentarismo-moody/sentiment"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestPtTokenize(t *testing.T) {
	tokenizer := sentiment.NewPtTokenizer()
	text := "E vida de gado, povo marcado é povo feliz!!!"
	m := tokenizer.Tokenize(text)
	log.Printf("%v", m)

	testWords := map[string]int64{
		"liz":1,
		"vida":1,
		"gad":1,
		"marcad":1,
	}

	for testWord, testWordValue  := range testWords {
		valid := false
		for foundWord, foundWordValue  := range m {
			if testWord == foundWord {
				//testWordValue
				//log.Println(testWord,testWordValue,foundWord,foundWordValue)
				assert.Equal(t, testWordValue, foundWordValue)
				valid = true
			}
		}
		if !valid {
			t.Fatal("Word expected never found --> ",testWord)
		}
	}

}
