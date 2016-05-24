package model

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/eaigner/shield"
	"log"
	"os"
)

var REDIS_HOST = os.Getenv("REDIS_HOST")
var REDIS_PORT = os.Getenv("REDIS_PORT")
var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

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
		REDIS_HOST = "g7-box"
	}
	if REDIS_PORT == "" {
		REDIS_HOST = "6379"
	}
	if REDIS_PASSWORD == "" {
	}
	if shieldInstance == nil {
		shieldInstance = shield.New(
			shield.NewEnglishTokenizer(),
			shield.NewRedisStore(REDIS_HOST+":"+REDIS_PORT, REDIS_PASSWORD, log.New(&buf, "logger: ", log.Lshortfile), ""),
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

var sentimentList map[string]string

func init(){
	sentimentList = make(map[string]string)
	sentimentList["-5"] = "Terrible!"
	sentimentList["-4"] = "Sucks"
	sentimentList["-3"] = "Bad"
	sentimentList["-2"] = "Not Good"
	sentimentList["-1"] = "Eh"
	sentimentList["0"] = "Neutral"
	sentimentList["1"] = "OK"
	sentimentList["2"] = "Good"
	sentimentList["3"] = "Like It"
	sentimentList["4"] = "Loved It"
	sentimentList["5"] = "Awesome!"
	sentimentList["6"] = "Unknown"
}
// GetSentiment classifies a single string of text. Returns the tag it matched.
func GetSentiment(text string) string {
	InitShield()
	tag, err := shieldInstance.Classify(text)
	if err != nil || tag == "" {
		tag = "6"
	}
	//log.Println(text," classified as ", sentimentList[tag])
	return sentimentList[tag]
}
