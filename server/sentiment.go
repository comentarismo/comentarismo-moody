package server

import (
	model "comentarismo-moody/model"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"time"
	"os"
	"strings"
)

var (
	FACEBOOK_KEY = os.Getenv("FACEBOOK_KEY")
	FACEBOOK_SECRET = os.Getenv("FACEBOOK_SECRET")

	YOUTUBE_KEY = os.Getenv("YOUTUBE_KEY")
	YOUTUBE_SECRET = os.Getenv("YOUTUBE_SECRET")

	INSTAGRAM_KEY = os.Getenv("INSTAGRAM_KEY")
	INSTAGRAM_SECRET = os.Getenv("INSTAGRAM_SECRET")

	VINEVIDEO_KEY = os.Getenv("VINEVIDEO_KEY")
	VINEVIDEO_SECRET = os.Getenv("VINEVIDEO_SECRET")

	Training = os.Getenv("training")
	IsTrained = false
)

func init() {

	if(YOUTUBE_KEY == ""){
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


func LoadTrainingData() bool {
	if Training == "" {
		Training = "./static/training/afinn-111.csv"
	}
	if IsTrained == false {
		log.Println("Training " + Training)
		trainingFiles := strings.Split(Training, ",")
		for _, path := range trainingFiles {
			model.LoadTrainingData(path)
		}
		IsTrained = true
	}
	return IsTrained
}


func RunReport(postURL string) []byte {
	// Parse URL
	notTrained := LoadTrainingData()
	if !notTrained {
		return jsonError("Unable to train the engine for this url.")
	}
	domain, urlParts := parseURL(postURL)
	if domain == "" || len(urlParts) == 0 {
		return jsonError("Unable to parse post url.")
	}

	// Create Report
	theReport := model.Report{URL: postURL}
	//var thePost Post

	log.Println("Going to run report for ", domain)

	provider, err := model.GetProvider(domain)
	if err != nil {
		log.Println(err)
		return jsonError("Could not GetProvider.")
	}

	err = provider.SetID(urlParts)
	if err != nil {
		//log.Println(err)
		log.Println("Could not SetID.")
		return jsonError("Could not SetID.")
	}

	//Fetch the metadata
	flag := provider.GetMetadata()

	if !flag {
		log.Println("Could not fetch metadata.")
		return jsonError("Could not fetch metadata.")
	}



	// Fetch the comments
	comments := provider.GetComments()

	//// If we don't get an comments back, wait for the metadata call to return and send an error.
	if comments.IsEmpty() {
		log.Println("No comments found for this post.")
		return jsonError("No comments found for this post.")
	}

	provider.SetReport(&theReport, comments)

	// Set comments returned
	theReport.CollectedComments = comments.GetTotal()
	theReport.CommentCoveragePercent = math.Ceil((float64(theReport.CollectedComments) / float64(theReport.TotalComments)) * float64(100))

	done := make(chan bool)

	// Set Keywords
	go func() {
		theReport.Keywords = comments.GetKeywords()
		done <- true
	}()

	// Sentiment Tagging
	go func() {
		theReport.Sentiment = comments.GetSentimentSummary()
		done <- true
	}()

	// Wait for everything to finish up
	for i := 0; i < 2; i++ {
		<-done
	}

	// Pull a few sample comments
	theReport.SampleComments = comments.GetRandom(3)
	theReport.SentimentList = comments.GetSentimentList()

	// Calculate Average Daily Comments
	timestamp, _ := strconv.ParseInt(theReport.PublishedAt, 10, 64)
	t := time.Unix(timestamp, 0)
	delta := time.Now().Sub(t)
	theReport.CommentAvgPerDay = float64(theReport.TotalComments) / (float64(delta.Hours()) / float64(24))

	reportJSON, err := json.Marshal(theReport)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("OK Going to return report for ", domain," TotalComments: ",theReport.TotalComments)

	log.Println("SentimentList len ",len(theReport.SentimentList))

	// Output Report
	return reportJSON
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
