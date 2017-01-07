package sentiment_test

import (
	"comentarismo-moody/sentiment"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestRuTokenize(t *testing.T) {
	tokenizer := sentiment.NewRuTokenizer()
	text := "Уже тошнит от таких новостей. Лучше бы деньги потраченные на заработную плату этих борцов с экстремизмом израсходовали на лечение больных детей, или сиротам еды бы приличной купили."
	m := tokenizer.Tokenize(text)
	log.Printf("%v", m)

	//map[чш:1 зм:1 зр:1 тр:2 нны:1 пл:1 рц:1 льных:1 ды:1]
	testWords := map[string]int64{
		"чш":    1,
		"зм":    1,
		"зр":    1,
		"тр":    2,
		"нны":   1,
		"пл":    1,
		"рц":    1,
		"льных": 1,
		"ды":    1,
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
