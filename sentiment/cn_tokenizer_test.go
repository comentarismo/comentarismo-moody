package sentiment_test

import (
	"comentarismo-moody/sentiment"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// go test  sentiment/cn_tokenizer_test.go -v
func TestCNTokenize(t *testing.T) {
	tokenizer := sentiment.NewCNTokenizer()
	text := "南京现在也有很多人加入挖比特币行业，刚刚兴起的，都靠偷国家的电赚钱"
	m := tokenizer.Tokenize(text)
	log.Printf("INFO: %v", m)

	//特币行业:1 刚刚兴:1 偷国家:1 电赚钱:1 南京现:1 加入挖:1
	testWords := map[string]int64{
		"特币行业": 1,
		"刚刚兴":  1,
		"偷国家":  1,
		"电赚钱":  1,
		"南京现":  1,
		"加入挖":   1,
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
