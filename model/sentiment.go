package model

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/eaigner/shield"
	"log"
	"os"
)

var REDIS_HOST = os.Getenv("redis_host")
var REDIS_PASS = os.Getenv("redis_pass")

var buf bytes.Buffer
var shieldInstance shield.Shield

// SentimentTag is a list entry of the tag and the percent of comments that were classified with that tag.
type SentimentTag struct {
	Name    string
	Percent float64
}

// InitShield instantiates the text classifier engine
func InitShield() {
	if REDIS_HOST == "" {
		REDIS_HOST = "192.168.0.42:6379"
	}
	if REDIS_PASS == "" {
	}
	if shieldInstance == nil {
		shieldInstance = shield.New(
			shield.NewEnglishTokenizer(),
			shield.NewRedisStore(REDIS_HOST, REDIS_PASS, log.New(&buf, "logger: ", log.Lshortfile), ""),
		)
	}
}

// LoadTrainingData input the training data in to the text classifier
func LoadTrainingData(path string) {
	csvfile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	InitShield()

	fmt.Println("Learning started.")

	for _, row := range csvData {
		// score, _ := strconv.ParseInt(row[1], 10, 0)
		shieldInstance.Learn(row[1], row[0])
	}

	fmt.Println("Learning complete!")
}

// GetSentiment classifies a single string of text. Returns the tag it matched.
func GetSentiment(text string) string {
	InitShield()

	tag, err := shieldInstance.Classify(text)
	if err == nil {
		return tag
	}

	return "UNKNOWN"
}
