package server

import (
	"comentarismo-moody/model"
	"encoding/json"
	"errors"
	r "github.com/dancannon/gorethink"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	FACEBOOK_KEY    = os.Getenv("FACEBOOK_KEY")
	FACEBOOK_SECRET = os.Getenv("FACEBOOK_SECRET")

	YOUTUBE_KEY    = os.Getenv("YOUTUBE_KEY")
	YOUTUBE_SECRET = os.Getenv("YOUTUBE_SECRET")

	INSTAGRAM_KEY    = os.Getenv("INSTAGRAM_KEY")
	INSTAGRAM_SECRET = os.Getenv("INSTAGRAM_SECRET")

	VINEVIDEO_KEY    = os.Getenv("VINEVIDEO_KEY")
	VINEVIDEO_SECRET = os.Getenv("VINEVIDEO_SECRET")

	IsTrainedEn = false
	IsTrainedEs = false
	IsTrainedPt = false
	IsTrainedIt = false
	IsTrainedFr = false
	IsTrainedHr = false
	IsTrainedRu = false
)

func init() {

	if YOUTUBE_KEY == "" {
		YOUTUBE_KEY = "AIzaSyBameiyxxJw0W27lydpPuPRocfvGza9gXM"
	}

	if FACEBOOK_KEY == "" {
		FACEBOOK_KEY = "118792195142376"
	}
	if FACEBOOK_SECRET == "" {
		FACEBOOK_SECRET = "49e8c46b9bec4f484fa045cedac63ea2"
	}

}

type WebError struct {
	Error string
}

var mu sync.Mutex

func LoadTrainingData(lang string) (IsTrained bool) {
	mu.Lock()
	pwd, _ := os.Getwd()

	targetDir := pwd + "/static/training/afinn-111-en.csv"
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		targetDir = pwd + "/../static/training/afinn-111-en.csv"
	}
	if lang == "pt" {
		targetDir = pwd + "/static/training/afinn-111-pt.csv"
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			targetDir = pwd + "/../static/training/afinn-111-pt.csv"
		}
	} else if lang == "es" {
		targetDir = pwd + "/static/training/afinn-111-es.csv"
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			targetDir = pwd + "/../static/training/afinn-111-es.csv"
		}
	} else if lang == "fr" {
		targetDir = pwd + "/static/training/afinn-111-fr.csv"
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			targetDir = pwd + "/../static/training/afinn-111-fr.csv"
		}
	} else if lang == "it" {
		targetDir = pwd + "/static/training/afinn-111-it.csv"
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			targetDir = pwd + "/../static/training/afinn-111-it.csv"
		}
	} else if lang == "hr" {
		targetDir = pwd + "/static/training/afinn-111-hr.csv"
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			targetDir = pwd + "/../static/training/afinn-111-hr.csv"
		}
	} else if lang == "ru" {
		targetDir = pwd + "/static/training/afinn-111-ru.csv"
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			targetDir = pwd + "/../static/training/afinn-111-ru.csv"
		}
	} else {
		//panic("Language is not supported, "+lang)
		//log.Println("Will use default lang en")
		lang = "en"
	}

	//check with redis if we are already trained
	//istrained, err := Client.Get("IsTrained").Result()
	//if err == redis.Nil {
	//	IsTrained = false
	//} else if err != nil {
	//	IsTrained = false
	//}else if istrained == "true" {
	//	IsTrained = true
	//}

	IsTrained = IsLangTrained(lang)

	if IsTrained == false {

		log.Println("Training " + targetDir)
		trainingFiles := strings.Split(targetDir, ",")
		for _, path := range trainingFiles {

			err := model.LoadTrainingData(lang, path)
			if err != nil {

			}
		}
		IsTrained = true
		SetLangTrained(lang)
		////set redis
		//Client.Set("IsTrained", "true", time.Minute * 1)
		//Client.Set("IsTrained", "true", time.Hour * 24)

	}
	mu.Unlock()

	return
}

func SetLangTrained(lang string) {
	if lang == "en" {
		IsTrainedEn = true
	} else if lang == "pt" {
		IsTrainedPt = true
	} else if lang == "es" {
		IsTrainedEs = true
	} else if lang == "fr" {
		IsTrainedFr = true
	} else if lang == "it" {
		IsTrainedIt = true
	} else if lang == "hr" {
		IsTrainedHr = true
	} else if lang == "ru" {
		IsTrainedRu = true
	}
}

func IsLangTrained(lang string) (isTrained bool) {
	if lang == "en" {
		isTrained = IsTrainedEn
	} else if lang == "pt" {
		isTrained = IsTrainedPt
	} else if lang == "es" {
		isTrained = IsTrainedEs
	} else if lang == "fr" {
		isTrained = IsTrainedFr
	} else if lang == "it" {
		isTrained = IsTrainedIt
	} else if lang == "hr" {
		isTrained = IsTrainedHr
	} else if lang == "ru" {
		isTrained = IsTrainedRu
	}
	return
}

func RunQuickReport(postURL, lang string) (model.Report, error) {
	domain, urlParts := parseURL(postURL)
	if domain == "" || len(urlParts) == 0 {
		return model.Report{}, errors.New("RunQuickReport, Unable to parse post url.")
	}

	// Create Report
	theReport := model.Report{URL: postURL}
	//var thePost Post

	log.Println("RunQuickReport, Going to run report for ", domain)

	provider, err := model.GetProvider(domain)
	if err != nil {
		log.Println("RunQuickReport, Could not GetProvider.", err)
		return model.Report{}, errors.New("RunQuickReport, Could not GetProvider.")
	}

	err = provider.SetID(urlParts)
	if err != nil {
		//log.Println(err)
		log.Println("RunQuickReportm, Could not SetID.")
		return model.Report{}, errors.New("RunQuickReport, Could not SetID.")
	}

	err = provider.SetLang(lang)
	if err != nil {
		//log.Println(err)
		log.Println("RunQuickReport, Could not LANG on provider.")
		return model.Report{}, errors.New("RunQuickReport, Could not LANG on provider.")
	}

	//Fetch the metadata
	flag := provider.GetMetadata()
	if !flag {
		log.Println("RunQuickReport, Could not fetch metadata.")
		return model.Report{}, errors.New("RunQuickReport, Could not fetch metadata.")
	}

	provider.SetReport(&theReport, model.CommentList{})

	// Set comments returned
	theReport.CollectedComments = 0
	theReport.CommentCoveragePercent = 0

	//set date
	theReport.Date = time.Now()

	// Calculate Average Daily Comments
	timestamp, _ := strconv.ParseInt(theReport.PublishedAt, 10, 64)
	t := time.Unix(timestamp, 0)
	delta := time.Now().Sub(t)
	theReport.CommentAvgPerDay = float64(theReport.TotalComments) / (float64(delta.Hours()) / float64(24))

	log.Println("RunQuickReport, OK Going to return report for ", theReport)

	// Output Report
	return theReport, nil
}

func RunReport(postURL, lang string) (model.Report, error) {
	// Parse URL
	notTrained := LoadTrainingData(lang)
	if !notTrained {
		log.Println("Unable to train the engine for this url.")
		return model.Report{}, errors.New("Unable to train the engine for this url.")

	}
	domain, urlParts := parseURL(postURL)
	if domain == "" || len(urlParts) == 0 {
		return model.Report{}, errors.New("Unable to parse post url.")
	}

	// Create Report
	theReport := model.Report{URL: postURL}
	//var thePost Post

	log.Println("Going to run report for ", domain)

	provider, err := model.GetProvider(domain)
	if err != nil {
		log.Println(err)
		return model.Report{}, errors.New("Could not GetProvider.")
	}

	err = provider.SetID(urlParts)
	if err != nil {
		//log.Println(err)
		log.Println("Could not SetID.")
		return model.Report{}, errors.New("Could not SetID.")
	}

	err = provider.SetLang(lang)
	if err != nil {
		//log.Println(err)
		log.Println("Could not LANG on provider.")
		return model.Report{}, errors.New("Could not LANG on provider.")
	}

	//Fetch the metadata
	flag := provider.GetMetadata()

	if !flag {
		log.Println("Could not fetch metadata.")
		return model.Report{}, errors.New("Could not fetch metadata.")
	}

	// Fetch the comments
	resultsChannel := make(chan *model.Comment)
	countChannel := make(chan int)
	var totalCount int
	var finished bool
	ct := 1
	go provider.GetCommentsChan(resultsChannel, countChannel)
	go func() {
		for range time.Tick(time.Minute * time.Duration(ct)) {
			log.Printf("Total comments processed last %d minute is %d, per sec: %d", ct, totalCount, totalCount/(ct*60))
			totalCount = 0
			if finished {
				break
			}
		}
	}()
	go func() {
		for c := range countChannel {
			if c == 0 {
				finished = true
				break
			}
			totalCount += c
		}
	}()

	totalProcessed := 0
	var commentsList model.CommentList
	_uuid, err := UUID(postURL)

	if err != nil {
		log.Println("Error: RunReport, UUID, ", err)
	}

	for comment := range resultsChannel {
		totalProcessed += 1

		comment.GetSentiment()

		//save to DB
		comment.UUID = _uuid
		//log.Println("RunReport, save comment into database")
		_, err = r.Table("commentaries_sentiment_report").Insert(comment, r.InsertOpts{Conflict: "update"}).RunWrite(Session)
		if err != nil {
			log.Println("ERROR: RunReport, INSERT sentiment TABLE JUST FAILED :|", err)
		}

		commentsList.Comments = append(commentsList.Comments, comment)

		if finished {
			break
		}
	}

	//// If we don't get an comments back, wait for the metadata call to return and send an error.
	if commentsList.IsEmpty() {
		log.Println("No comments found for this post.")
		return model.Report{}, errors.New("No comments found for this post.")
	}
	provider.SetReport(&theReport, commentsList)

	// Set comments returned
	theReport.CollectedComments = commentsList.GetTotal()
	theReport.CommentCoveragePercent = math.Ceil((float64(theReport.CollectedComments) / float64(theReport.TotalComments)) * float64(100))

	//set date
	theReport.Date = time.Now()

	done := make(chan bool)

	// Set Keywords
	go func() {
		theReport.Keywords = commentsList.GetKeywords()
		done <- true
	}()

	// Sentiment Tagging
	go func() {
		theReport.Sentiment = commentsList.GetSentimentSummary()
		done <- true
	}()

	// Sentiment Scores
	go func() {
		theReport.SentimentScores = commentsList.GetSentimentScores()
		done <- true
	}()

	// Wait for everything to finish up
	for i := 0; i < 3; i++ {
		<-done
	}

	// Pull a few sample comments
	//theReport.SampleComments = comments.GetRandom(3)
	theReport.SentimentList = commentsList.GetSentimentList()

	// Calculate Average Daily Comments
	timestamp, _ := strconv.ParseInt(theReport.PublishedAt, 10, 64)
	t := time.Unix(timestamp, 0)
	delta := time.Now().Sub(t)
	theReport.CommentAvgPerDay = float64(theReport.TotalComments) / (float64(delta.Hours()) / float64(24))

	log.Println("OK Going to return report for ", domain, " TotalComments: ", theReport.TotalComments)

	log.Println("SentimentList len ", len(theReport.SentimentList))

	// Output Report
	return theReport, nil
}

func jsonError(msg string) []byte {
	errorJSON, _ := json.Marshal(WebError{Error: msg})
	return errorJSON
}

func parseURL(url string) (string, []string) {
	sites := map[string]string{
		"instagram": "instag\\.?ram(\\.com)?/p/([\\w]*)/?",
		"youtube":   "youtu\\.?be(\\.?com)?/(watch\\?v=)?([\\w\\-_]*)",
		"facebook":  "facebook\\.com/([\\w]*)/posts/([\\d]*)/?",
		"vine":      "vine\\.co/v/([\\w]*)?",
	}

	var domain string
	var matches []string

	for d, rstr := range sites {
		r, _ := regexp.Compile(rstr)
		matches = r.FindStringSubmatch(url)
		if len(matches) > 0 {
			domain = d
			break
		}
	}

	return domain, matches
}
