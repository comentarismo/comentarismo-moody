package sentiment_test

import (
	"comentarismo-moody/sentiment"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestHrTokenize(t *testing.T) {
	tokenizer := sentiment.NewHrTokenizer()
	text := "Ako vam je ovo picka ??? Kao da se vratila iz Auschwitza sa godisnjeg odmora"
	m := tokenizer.Tokenize(text)
	log.Printf("%v", m)

	//map[ak:1 ck:1 vr:1 chw:1 dm:1]
	testWords := map[string]int64{
		"ak":  1,
		"ck":  1,
		"vr":  1,
		"chw": 1,
		"dm":  1,
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
