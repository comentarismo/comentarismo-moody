package sentiment_test

import (
	"testing"
	"comentarismo-moody/sentiment"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestEsTokenize(t *testing.T) {
	tokenizer := sentiment.NewEsTokenizer()
	text := "Por el amor de Dios!! En que quedamos!?"
	m := tokenizer.Tokenize(text)

	testWords := map[string]int64{
		"por":1,
		"amor":1,
		"dios":1,
		"que":1,
		"quedamos":1,
	}

	//log.Println(m)

	for testWord,testWordValue  := range testWords {
		valid := false
		for foundWord, foundWordValue  := range m {
			if testWord == foundWord {
				//log.Println(testWord, testWordValue,foundWord, foundWordValue)
				assert.Equal(t,testWordValue,foundWordValue)
				valid = true
			}
		}
		if !valid {
			t.Fatal("Word expected never found --> ",testWord)
		}
	}

}
