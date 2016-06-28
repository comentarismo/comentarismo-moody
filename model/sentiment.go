package model

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"comentarismo-moody/sentiment"
)

var REDIS_HOST = os.Getenv("REDIS_HOST")
var REDIS_PORT = os.Getenv("REDIS_PORT")
var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

var buf bytes.Buffer
var sentimentInstanceEn sentiment.Sentiment
var sentimentInstancePt sentiment.Sentiment
var sentimentInstanceEs sentiment.Sentiment
var sentimentInstanceFr sentiment.Sentiment
var sentimentInstanceIt sentiment.Sentiment
var sentimentInstanceHr sentiment.Sentiment
var sentimentInstanceRu sentiment.Sentiment

// InitShield instantiates the text classifier engine
func InitSentiment(lang string) {
	if REDIS_HOST == "" {
		REDIS_HOST = "g7-box"
	}
	if REDIS_PORT == "" {
		REDIS_PORT = "6379"
	}
	if REDIS_PASSWORD == "" {
	}

	//Get tokenizers
	var tokenizer sentiment.Tokenizer
	tokenizer = GetTokenizerForLang(lang)



	if lang == "en" && sentimentInstanceEn == nil {
		//create alert msg
		log.Println("Starting redis "+lang+" text classifier engine, ",REDIS_HOST+":"+REDIS_PORT)
		//create redis instance
		store := sentiment.NewRedisStore(REDIS_HOST+":"+REDIS_PORT, REDIS_PASSWORD, log.New(&buf, "logger: ", log.Lshortfile), "")
		//start sentiment instance for lang
		sentimentInstanceEn = sentiment.New(
			tokenizer,
			store,
		)
	} else if lang == "pt" && sentimentInstancePt == nil {
		//create alert msg
		log.Println("Starting redis "+lang+" text classifier engine, ",REDIS_HOST+":"+REDIS_PORT)
		//create redis instance
		store := sentiment.NewRedisStore(REDIS_HOST+":"+REDIS_PORT, REDIS_PASSWORD, log.New(&buf, "logger: ", log.Lshortfile), "")
		//start sentiment instance for lang
		sentimentInstancePt = sentiment.New(
			tokenizer,
			store,
		)
	} else if lang == "es" && sentimentInstanceEs == nil {
		//create alert msg
		log.Println("Starting redis "+lang+" text classifier engine, ",REDIS_HOST+":"+REDIS_PORT)
		//create redis instance
		store := sentiment.NewRedisStore(REDIS_HOST+":"+REDIS_PORT, REDIS_PASSWORD, log.New(&buf, "logger: ", log.Lshortfile), "")
		//start sentiment instance for lang
		sentimentInstanceEs = sentiment.New(
			tokenizer,
			store,
		)
	} else if lang == "fr" && sentimentInstanceFr == nil {
		//create alert msg
		log.Println("Starting redis "+lang+" text classifier engine, ",REDIS_HOST+":"+REDIS_PORT)
		//create redis instance
		store := sentiment.NewRedisStore(REDIS_HOST+":"+REDIS_PORT, REDIS_PASSWORD, log.New(&buf, "logger: ", log.Lshortfile), "")
		//start sentiment instance for lang
		sentimentInstanceFr = sentiment.New(
			tokenizer,
			store,
		)
	} else if lang == "it" && sentimentInstanceIt == nil {
		//create alert msg
		log.Println("Starting redis "+lang+" text classifier engine, ",REDIS_HOST+":"+REDIS_PORT)
		//create redis instance
		store := sentiment.NewRedisStore(REDIS_HOST+":"+REDIS_PORT, REDIS_PASSWORD, log.New(&buf, "logger: ", log.Lshortfile), "")
		//start sentiment instance for lang
		sentimentInstanceIt = sentiment.New(
			tokenizer,
			store,
		)
	}else if lang == "hr" && sentimentInstanceHr == nil {
		//create alert msg
		log.Println("Starting redis "+lang+" text classifier engine, ",REDIS_HOST+":"+REDIS_PORT)
		//create redis instance
		store := sentiment.NewRedisStore(REDIS_HOST+":"+REDIS_PORT, REDIS_PASSWORD, log.New(&buf, "logger: ", log.Lshortfile), "")
		//start sentiment instance for lang
		sentimentInstanceHr = sentiment.New(
			tokenizer,
			store,
		)
	}else if lang == "ru" && sentimentInstanceRu == nil {
		//create alert msg
		log.Println("Starting redis "+lang+" text classifier engine, ",REDIS_HOST+":"+REDIS_PORT)
		//create redis instance
		store := sentiment.NewRedisStore(REDIS_HOST+":"+REDIS_PORT, REDIS_PASSWORD, log.New(&buf, "logger: ", log.Lshortfile), "")
		//start sentiment instance for lang
		sentimentInstanceRu = sentiment.New(
			tokenizer,
			store,
		)
	}
}

func GetTokenizerForLang(lang string) (tokenizer sentiment.Tokenizer) {
	if lang == "en" && sentimentInstanceEn == nil {
		tokenizer = sentiment.NewTokenizer()
	} else if lang == "pt" && sentimentInstancePt == nil {
		tokenizer = sentiment.NewPtTokenizer()
	} else if lang == "es" && sentimentInstanceEs == nil {
		tokenizer = sentiment.NewEsTokenizer()
	} else if lang == "fr" && sentimentInstanceFr == nil {
		tokenizer = sentiment.NewFrTokenizer()
	} else if lang == "it" && sentimentInstanceIt == nil {
		tokenizer = sentiment.NewItTokenizer()
	} else if lang == "hr" && sentimentInstanceHr == nil {
		tokenizer = sentiment.NewHrTokenizer()
	} else if lang == "ru" && sentimentInstanceRu == nil {
		tokenizer = sentiment.NewRuTokenizer()
	}
	return
}

func GetSentimentInstanceForLang(lang string) (currentSentimentInstance sentiment.Sentiment){
	if lang == "en" {
		currentSentimentInstance = sentimentInstanceEn
	} else if lang == "pt" {
		currentSentimentInstance = sentimentInstancePt
	} else if lang == "es" {
		currentSentimentInstance = sentimentInstanceEs
	} else if lang == "fr" {
		currentSentimentInstance = sentimentInstanceFr
	} else if lang == "it" {
		currentSentimentInstance = sentimentInstanceIt
	} else if lang == "hr" {
		currentSentimentInstance = sentimentInstanceHr
	} else if lang == "ru" {
		currentSentimentInstance = sentimentInstanceRu
	}
	return
}

// LoadTrainingData input the training data in to the text classifier
func LoadTrainingData(lang, path string) (err error) {
	csvfile, err := os.Open(path)
	if err != nil {
		log.Println("Error: LoadTrainingData, ",err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Println("Error: LoadTrainingData, ",err)
		return
	}

	InitSentiment(lang)

	log.Println("LoadTrainingData, Learning for lang "+lang+" started.")

	//sets := []shield.Set{}
	//for _, row := range csvData {
	//	//c := strings.SplitN(v, " ", 2)
	//	sets = append(sets, shield.Set{
	//		Class: row[0],
	//		Text:  row[1],
	//	})
	//}

	//shieldInstance.BulkLearn(sets)

	sentimentInstance := GetSentimentInstanceForLang(lang)

	for _, row := range csvData {
		// score, _ := strconv.ParseInt(row[1], 10, 0)
		//log.Println(row[1], row[0])
		sentimentInstance.Learn(row[1], row[0])
	}

	log.Println("LoadTrainingData, Learning for lang "+lang+" complete!")

	return
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
func GetSentiment(lang, text string) (sentimentTop string, scores map[string]float64, logScores map[string]map[string]int64) {
	InitSentiment(lang)

	//sanitize the text
	//text = strings.Replace(text, "?", " ", 1)
	sentimentInstance := GetSentimentInstanceForLang(lang)

	tag, scores, logScores, err := sentimentInstance.Classify(text)
	if err != nil {
		log.Println("Error: GetSentiment after Classify ",text,err)
	} else if tag == "" {
		tag = "6"
	}

	if sentimentList[tag] == "" {
		log.Println("Error: TAG -> ",tag)
		panic("Sentiment could not be defined, it maybe that the engine has old test data and is impacting the classify algorithm. hint: use FLUSHALL to cleanup database ")
	}
	//log.Println("tag: ",tag)
	//log.Println(text," classified as ", sentimentList[tag])
	sentimentTop = sentimentList[tag]

	return
}
