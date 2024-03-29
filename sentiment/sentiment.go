package sentiment

import (
	"log"
	"math"
	"os"
	"sync"
)

const defaultProb float64 = 1e-11

var SENTIMENT_DEBUG = os.Getenv("SENTIMENT_DEBUG")

type TokenStore struct {
	mu        sync.Mutex
	tokenizer Tokenizer
	Store     Store
}

func New(t Tokenizer, s Store) Sentiment {
	return &TokenStore{
		tokenizer: t,
		Store:     s,
	}
}

func (sh *TokenStore) Learn(class, text string) (err error) {
	return sh.increment(class, text, 1)
}

func (sh *TokenStore) BulkLearn(sets []Set) (err error) {
	return sh.bulkIncrement(sets, 1)
}

func (sh *TokenStore) Forget(class, text string) (err error) {
	return sh.increment(class, text, -1)
}

func (sh *TokenStore) increment(class, text string, sign int64) (err error) {
	if len(class) == 0 {
		panic("invalid class")
	}
	if len(text) == 0 {
		panic("invalid text")
	}
	return sh.bulkIncrement([]Set{{Class: class, Text: text}}, sign)
}

func (sh *TokenStore) bulkIncrement(sets []Set, sign int64) (err error) {
	sh.mu.Lock()
	if len(sets) == 0 {
		sh.mu.Unlock()
		panic("invalid data set")
	}
	m := make(map[string]map[string]int64)
	for _, set := range sets {
		tokens := sh.tokenizer.Tokenize(set.Text)
		for k := range tokens {
			tokens[k] *= sign
		}
		if w, ok := m[set.Class]; ok {
			for word, count := range tokens {
				w[word] += count
			}
		} else {
			m[set.Class] = tokens
		}
	}
	for class := range m {
		if err = sh.Store.AddClass(class); err != nil {
			Debug("bulkIncrement, sh.Store.AddClass, ", err)
			sh.mu.Unlock()
			return
		}
	}

	err = sh.Store.IncrementClassWordCounts(m)
	sh.mu.Unlock()

	return
}

func getKeys(m map[string]int64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (s *TokenStore) Score(text string) (scores map[string]float64, classFreqs map[string]map[string]int64, err error) {
	s.mu.Lock()

	// Get total class word counts
	totals, err := s.Store.TotalClassWordCounts()
	if err != nil {
		s.mu.Unlock()
		return
	}

	Debug("shield.Score, totals", totals)
	classes := getKeys(totals)

	// Tokenize text
	wordFreqs := s.tokenizer.Tokenize(text)
	words := getKeys(wordFreqs)

	Debug("sentiment.Score, text ", text)
	Debug("sentiment.Score, wordFreqs ", wordFreqs)
	Debug("sentiment.Score, wordFreqs len ", len(wordFreqs))
	Debug("sentiment.Score, words", words)
	Debug("sentiment.Score, words len ", len(words))

	if len(wordFreqs) == 0 || len(words) == 0 {
		Debug("Error: We are not able to get wordFreqs or words, in order to prevent wrong number of arguments for 'hmget' command we will return now! ", len(words))
		s.mu.Unlock()
		return
	}

	// Get word frequencies for each class
	classFreqs = make(map[string]map[string]int64)
	for _, class := range classes {
		freqs, err2 := s.Store.ClassWordCounts(class, words)
		if err2 != nil {
			err = err2
			Debug("Error: shield.Score after ClassWordCounts -> ", err2)
			s.mu.Unlock()
			return
		}
		classFreqs[class] = freqs
	}

	s.mu.Unlock()

	// Calculate log scores for each class
	logScores := make(map[string]float64, len(classes))

	for _, class := range classes {
		freqs := classFreqs[class]
		total := totals[class]

		// Because this classifier is not biased, we don't use prior probabilities
		score := float64(0)
		for _, word := range words {
			// Compute the probability that this word belongs to that class
			wordProb := float64(freqs[word]) / float64(total)
			if wordProb == 0 {
				wordProb = defaultProb
			}
			score += math.Log(wordProb)
		}
		logScores[class] = score
	}

	// Normalize the scores
	var min = math.MaxFloat64
	var max = -math.MaxFloat64
	for _, score := range logScores {
		if score > max {
			max = score
		}
		if score < min {
			min = score
		}
	}
	r := max - min
	scores = make(map[string]float64, len(classes))
	for class, score := range logScores {
		if r == 0 {
			scores[class] = 1
		} else {
			scores[class] = (score - min) / r
		}
	}
	return
}

func (s *TokenStore) Classify(text string) (class string, scores map[string]float64, logScores map[string]map[string]int64, err error) {
	scores, logScores, err = s.Score(text)
	if err != nil {
		Debug("Error: shield.Classify-> ", err)
		return
	}

	// Select class with highest prob
	var score float64
	for k, v := range scores {
		if v > score {
			class, score = k, v
		}
	}
	Debug("class, score -> ", class, score)
	return
}

func (sh *TokenStore) Reset() error {
	return sh.Store.Reset()
}

func Debug(v ...interface{}) {
	if SENTIMENT_DEBUG == "true" {
		log.Println(v)
	}
}
