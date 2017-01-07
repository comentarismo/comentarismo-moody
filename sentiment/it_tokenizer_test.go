package sentiment_test

import (
	"comentarismo-moody/sentiment"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestItTokenize(t *testing.T) {
	tokenizer := sentiment.NewItTokenizer()
	text := "vista la vostra assolutà inutilità avreste fatto meglio a rinunciare del tutto allo stipendio"
	m := tokenizer.Tokenize(text)
	log.Printf("%v", m)

	testWords := map[string]int64{
		"vist": 1,
		"lut":  1,
		"inut": 1,
		"rin":  1,
		"tutt": 1,
		"stip": 1,
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
