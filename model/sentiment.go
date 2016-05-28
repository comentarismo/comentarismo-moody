package model

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"comentarismo-moody/sentiment"
)

var REDIS_HOST = os.Getenv("REDIS_HOST")
var REDIS_PORT = os.Getenv("REDIS_PORT")
var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

var buf bytes.Buffer
var sentimentInstance sentiment.Sentiment

// InitShield instantiates the text classifier engine
func InitSentiment() {
	if REDIS_HOST == "" {
		REDIS_HOST = "g7-box"
	}
	if REDIS_PORT == "" {
		REDIS_PORT = "6379"
	}
	if REDIS_PASSWORD == "" {
	}
	if sentimentInstance == nil {
		log.Println("Starting redis text classifier engine, ",REDIS_HOST+":"+REDIS_PORT)
		sentimentInstance = sentiment.New(
			sentiment.NewTokenizer(),
			sentiment.NewRedisStore(REDIS_HOST+":"+REDIS_PORT, REDIS_PASSWORD, log.New(&buf, "logger: ", log.Lshortfile), ""),
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

	InitSentiment()

	fmt.Println("Learning started.")

	//sets := []shield.Set{}
	//for _, row := range csvData {
	//	//c := strings.SplitN(v, " ", 2)
	//	sets = append(sets, shield.Set{
	//		Class: row[0],
	//		Text:  row[1],
	//	})
	//}

	//shieldInstance.BulkLearn(sets)

	for _, row := range csvData {
		// score, _ := strconv.ParseInt(row[1], 10, 0)
		//log.Println(row[1], row[0])
		sentimentInstance.Learn(row[1], row[0])
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
	InitSentiment()

	//sanitize the text
	//text = strings.Replace(text, "?", " ", 1)

	tag, err := sentimentInstance.Classify(text)
	if err != nil {
		log.Println("Error: GetSentiment after Classify ",text,err)
	} else if tag == "" {
		tag = "6"
	}

	if sentimentList[tag] == "" {
		panic("Sentiment could not be defined, it maybe that the engine has old test data and is impacting the classify algorithm. hint: use FLUSHALL to cleanup database ")
	}
	log.Println("tag: ",tag)
	log.Println(text," classified as ", sentimentList[tag])
	return sentimentList[tag]
}
