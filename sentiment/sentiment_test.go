package sentiment_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
	"comentarismo-moody/sentiment"
)

var REDIS_HOST = os.Getenv("REDIS_HOST")
var REDIS_PORT = os.Getenv("REDIS_PORT")
var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

func readDataSet(dataFile, labelFile string, t *testing.T) []string {
	d, err := ioutil.ReadFile("testdata/" + dataFile)
	if err != nil {
		t.Fatal(err)
	}
	l, err := ioutil.ReadFile("testdata/" + labelFile)
	if err != nil {
		t.Fatal(err)
	}
	dl := strings.Split(string(d), "\n")
	ll := strings.Split(string(l), "\n")
	x, y := len(dl), len(ll)
	if x != y {
		t.Fatal(x, y)
	}
	var a []string
	for i, v := range ll {
		k := strings.TrimSpace(v)
		if k != "" {
			a = append(a, fmt.Sprintf("%s %s", k, strings.TrimSpace(dl[i])))
		}
	}
	return a
}

func newSentiment() sentiment.Sentiment {
	if REDIS_HOST == "" {
		REDIS_HOST = "g7-box"
	}
	if REDIS_PORT == "" {
		REDIS_PORT = "6379"
	}
	if REDIS_PASSWORD == "" {
	}

	logger := log.New(os.Stderr, "", log.LstdFlags)
	store := sentiment.NewRedisStore(REDIS_HOST + ":" + REDIS_PORT, REDIS_PASSWORD, logger, "redis")
	tokenizer := sentiment.NewTokenizer()

	sh := sentiment.New(tokenizer, store)
	err := sh.Reset()
	if err != nil {
		panic(err)
	}
	return sh
}

func TestLearn(t *testing.T) {
	sh := newSentiment()
	testData := readDataSet("testdata.txt", "testlabels.txt", t)
	trainData := readDataSet("traindata.txt", "trainlabels.txt", t)

	// Run on test sets
	sets := []sentiment.Set{}
	for _, v := range trainData {
		c := strings.SplitN(v, " ", 2)
		sets = append(sets, sentiment.Set{
			Class: c[0],
			Text:  c[1],
		})
	}

	sh.BulkLearn(sets)

	var hit, miss int
	for _, v2 := range testData {
		c := strings.SplitN(v2, " ", 2)
		k, v := c[0], c[1]
		clz, err := sh.Classify(v)
		if err != nil {
			t.Fatal(err, "key:", k, "value:", v)
		}
		if clz != k {
			miss++
		} else {
			hit++
		}
	}

	// Test hit/miss ratio
	// TODO: Tweak this, where possible
	minHitRatio := 0.73
	hitRatio := (float64(hit) / float64(hit + miss))
	if hitRatio < minHitRatio {
		t.Fatalf("%d hits, %d misses (expected ratio %.2f, is %.2f)", hit, miss, minHitRatio, hitRatio)
	}
}

func TestDecrement(t *testing.T) {
	sh := newSentiment()
	sh.Learn("a", "hello")
	sh.Learn("a", "sunshine")
	sh.Learn("a", "tree")
	sh.Learn("a", "water")

	sh.Learn("b", "iamb!")

	sh.Forget("a", "hello tree")
	sh.Forget("a", "hello")

	s := sh.(*sentiment.TokenStore)
	testWords := []string {
		"hello",
		"sunshine",
		"tree",
		"water",
	}
	m, err := s.Store.ClassWordCounts("a", testWords)
	if err != nil {
		t.Fatal(err)
	}

	testWordsAssert := map[string]int64{
		"hello":0,
		"sunshine":1,
		"tree":0,
		"water":1,
	}

	for testWord,v  := range testWordsAssert {
		valid := false
		for foundWord, v2  := range m {
			if testWord == foundWord {
				log.Println(testWord,v,foundWord,v2)
				assert.Equal(t, v, v2)
				valid = true
			}
		}
		if !valid {
			t.Fatal("Word expected never found --> ",testWord)
		}
	}

	testWords = []string{
		"iamb!",
		"hello",
	}

	m2, err := s.Store.ClassWordCounts("b", testWords)
	if err != nil {
		t.Fatal("ClassWordCounts",err)
	}

	testWordsAssert = map[string]int64{
		"iamb!":0,
		"hello":0,
	}

	for testWord, v  := range testWordsAssert {
		valid := false
		for foundWord, v2  := range m2 {
			if testWord == foundWord {
				log.Println(testWord,v,foundWord,v2)
				assert.Equal(t, v, v2)

				valid = true
				break
			}
		}
		if !valid {
			t.Fatal("Word expected never found --> ",testWord)
		}
	}

	wc, err := s.Store.TotalClassWordCounts()
	if err != nil {
		t.Fatal("Error: TotalClassWordCounts",err)
	}
	if x := len(wc); x != 2 {
		t.Fatal("x := len(wc); x != 2, ",x)
	}
	if x := wc["a"]; x != 2 {
		t.Fatal("x := wc['a']; x != 2, ",x)
	}
	if x := wc["b"]; x != 1 {
		t.Fatal("x := wc['b']; x != 1, ",x)
	}
}
