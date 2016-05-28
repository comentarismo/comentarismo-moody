package sentiment_test

import (
	"testing"
	"comentarismo-moody/sentiment"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	tokenizer := sentiment.NewTokenizer()
	text := "lorem    ipsum able hello erik    can do hi there  \t  spaaace! lorem"
	m := tokenizer.Tokenize(text)

	testWords := map[string]int64{
		"lorem":2,
		"ipsum":1,
		"hello":1,
		"erik":1,
		"spaaace":1,
	}

	log.Println(m)

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
