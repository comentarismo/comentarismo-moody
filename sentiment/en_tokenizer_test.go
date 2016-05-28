package sentiment_test

import (
	"testing"
	"comentarismo-moody/sentiment"
)

func TestTokenize(t *testing.T) {
	tokenizer := sentiment.NewTokenizer()
	text := "lorem    ipsum able hello erik    can do hi there  \t  spaaace! lorem"
	m := tokenizer.Tokenize(text)

	testWords := []string{
		"lorem",
		"ipsum",
		"hello",
		"erik",
		"spaaace",
	}

	for _, testWord := range testWords {
		valid := false
		for foundWord,_  := range m {
			if testWord == foundWord {
				valid = true
			}
		}
		if !valid {
			t.Fatal("Word expected never found --> ",testWord)
		}
	}

}
