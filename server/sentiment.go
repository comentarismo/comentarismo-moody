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
)

// Post is the interface for all the various post types (YouTubeVideo, etc...)
//type Post interface {
//	GetComments() model.CommentList
//	GetMetadata() bool
//}

type WebError struct {
	Error string
}

func RunReport(postURL string) []byte {
	// Parse URL
	domain, urlParts := parseURL(postURL)
	if domain == "" || len(urlParts) == 0 {
		return jsonError("Unable to parse post url.")
	}

	// Create Report
	theReport := model.Report{URL: postURL}
	//var thePost Post

	log.Println("Going to run report for ", domain)

	//switch domain {
	//case "youtube":
	//	if model.YouTubeKey == "" {
	//		model.YouTubeKey = "AIzaSyBameiyxxJw0W27lydpPuPRocfvGza9gXM"
	//		//return jsonError("API key for YouTube not configured.")
	//	}
	//
	//	thePost = &model.YouTubeVideo{ID: urlParts[len(urlParts)-1]}
	//case "instagram":
	//	if model.InstagramKey == "" {
	//		return jsonError("API key for Instagram not configured.")
	//	}
	//
	//	thePost = &model.InstagramPic{ShortCode: urlParts[len(urlParts)-1]}
	//case "facebook":
	//	if facebook.FacebookKey == "" || facebook.FacebookSecret == "" {
	//		return jsonError("Missing Facebook API credentials.")
	//	}
	//
	//	thePost = &facebook.Provider{PageName: urlParts[len(urlParts)-2], ID: urlParts[len(urlParts)-1]}
	//case "vine":
	//	if model.VineUsername == "" || model.VinePassword == "" {
	//		return jsonError("Missing Vine user credentials.")
	//	}
	//
	//	thePost = &model.VineVideo{ShortCode: urlParts[len(urlParts)-1]}
	//}

	provider, err := model.GetProvider(domain)
	if err != nil {
		log.Println(err)
		return jsonError("Could not fetch metadata.")
	}

	err = provider.SetID(urlParts)
	if err != nil {
		//log.Println(err)
		log.Println("karai")
		return jsonError("Could not fetch metadata.")
	}

	 //Fetch the metadata
	flag := provider.GetMetadata()

	if !flag {
		return jsonError("Could not fetch metadata.")
	}

	// Fetch the comments
	comments := provider.GetComments()

	//log.Println(len(comments.Comments))
	//
	//// If we don't get an comments back, wait for the metadata call to return and send an error.
	if comments.IsEmpty() {
		return jsonError("No comments found for this post.")
	}

	provider.SetReport(&theReport, comments)
	//
	////case *model.InstagramPic:
	////	theReport.Type = "InstagramPic"
	////	theReport.ID = p.ID
	////	theReport.Title = p.Caption
	////	theReport.PublishedAt = p.PublishedAt
	////	theReport.TotalComments = p.TotalComments
	////	theReport.Metadata = p
	////case *model.VineVideo:
	////	theReport.Type = "VineVideo"
	////	theReport.ID = p.ID
	////	theReport.PublishedAt = p.PublishedAt
	////	theReport.TotalComments = p.TotalComments
	////	theReport.Metadata = p

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

	// Calculate Average Daily Comments
	timestamp, _ := strconv.ParseInt(theReport.PublishedAt, 10, 64)
	t := time.Unix(timestamp, 0)
	delta := time.Now().Sub(t)
	theReport.CommentAvgPerDay = float64(theReport.TotalComments) / (float64(delta.Hours()) / float64(24))

	reportJSON, err := json.Marshal(theReport)
	if err != nil {
		fmt.Println(err)
	}

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
