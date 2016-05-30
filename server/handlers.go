package server

import (
	//util "comentarismo-moody/util"
	"encoding/json"
	"log"
	"net/http"
	"gopkg.in/redis.v3"
	"time"
	"fmt"
	r "github.com/dancannon/gorethink"

	"comentarismo-moody/model"
)

func init() {

}

func MoodyHandler(w http.ResponseWriter, req *http.Request) {
	AllowOrigin(w, req)

	postURL := req.URL.Query().Get("vid")
	refresh := req.URL.Query().Get("refresh")
	lang := req.URL.Query().Get("lang")

	if lang == "" {
		lang = "en"
	}
	//validate inputs
	if postURL == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing post URL or LANG"})
		w.Write(jsonBytes)
		return
	}

	log.Println("ReportHandler ", postURL, refresh)
	var jsonBytes []byte

	if (refresh == "") {
		//get from redis if available
		//serve from redis
		valid := true

		cached, err := Client.Get("report_" + postURL).Result()
		if err == redis.Nil {
			log.Println("report_" + postURL + " does not exists")
		} else if err != nil {
			panic(err)
		} else {
			log.Println("Serve from cache")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cached))
			return
		}

		log.Println("not in redis, is it on rethinkdb ??")

		/** BEGIN Code to unmarshal the url into a valid ID that we can use in our internal database **/
		domain, urlParts := parseURL(postURL)
		if domain == "" || len(urlParts) == 0 {
			log.Println("parseURL -> report_" + postURL + " does not exists")
			log.Println("404 not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		provider, err := model.GetProvider(domain)
		if err != nil {
			log.Println("Could not GetProvider. report_" + postURL + " does not exists", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = provider.SetID(urlParts)
		if err != nil {
			log.Println("Could not SetID. report_" + postURL + " does not exists", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		theReport := model.Report{}
		provider.SetReport(&theReport, model.CommentList{})
		/** END Code to unmarshal the url into a valid ID that we can use in our internal database **/

		simpleReport := model.Report{}
		commentsReport := model.Report{}


		/** BEGIN recover the simple sentiment report **/
		res, err := r.Table("sentiment_report").Get(theReport.ID).Run(Session)

		//if found rethinkdb sentiment_report, query all comments and save cache redis
		if err != nil || res == nil {
			log.Println("Could not get report_" + postURL + " does not exists", err)
			valid = false
		} else if res.IsNil() {
			log.Println("Could not SetID. report_" + postURL + " does not exists", err)
			valid = false
		} else {
			defer res.Close()
			res.One(&simpleReport)
			if (simpleReport.ID == "") {
				valid = false
				log.Println("Could NOT GET ID. report_" + postURL + " exists")
			}
		}
		/** END recover the simple sentiment report **/

		/** BEGIN recover the comments sentiment report **/
		if valid {
			log.Println("load sentiment_report ok, going to load all comments now, this may break :| ")
			res2, err := r.Table("sentiment").Get(theReport.ID).Pluck("sentimentlist", "topcomments").Run(Session)
			if err != nil || res2 == nil {
				log.Println("Could not get report_" + postURL + " does not exists", err)
				valid = false
			} else if res2.IsNil() {
				log.Println("Could not SetID. report_" + postURL + " does not exists", err)
				valid = false
			} else {
				defer res2.Close()
				res2.One(&commentsReport)
			}
		}
		/** END recover the comments sentiment report **/


		if valid {

			log.Println("SentimentList: ", len(commentsReport.SentimentList))
			//log.Println("SampleComments: ",len(item.SampleComments))
			log.Println("TopComments: ", len(commentsReport.TopComments))

			if len(commentsReport.SentimentList) > 0 {

				//merge both into one

				simpleReport.SentimentList = commentsReport.SentimentList
				simpleReport.TopComments = commentsReport.TopComments

				jsonBytes, err = json.Marshal(&simpleReport)

				if err != nil {
					fmt.Println("Error: ", err)
				} else {

					log.Println("loading from rethinkdb went ok :D ")

					Client.Set("report_" + postURL, jsonBytes, time.Hour * 1)
					log.Println("save to redis ok")

					w.Header().Set("Content-Type", "application/json")
					w.Write(jsonBytes)
					return;
				}
			} else {
				log.Println("Error: OMG we could not get this SentimentList to work :| No dougnuts for you today o__0 ")
			}
		} else {
			log.Println("Error: OMG we could not get this validation to work :| No dougnuts for you today o__0 ")
		}

	}

	log.Println("Will generate a new report -->", postURL)

	theReport, err := RunReport(postURL,lang)
	if err != nil {
		log.Println("Error: MoodyHandler() RunReport() ",err)
		w.WriteHeader(http.StatusInternalServerError)
		jsonBytes, _ := json.Marshal(WebError{Error: "RunReport failed :| -> "})
		w.Write(jsonBytes)
		return
		//return here ???

	}
	//save to rethinkdb

	//create simple report
	simpleReport := CreateSimpleReport(theReport)

	//create comments report
	commentsReport := CreateCommentsReport(theReport)

	update := true

	res, err := r.Table("sentiment_report").Get(theReport.ID).Pluck("id").Run(Session)
	if err != nil {
		update = false
	} else if res.IsNil() {
		update = false
	} else {
		defer res.Close()
		item := model.Report{}
		res.One(&item)
		if (item.ID == "") {
			update = false
		}
		log.Println("lets update sentiment report")
		if item.ID == "" {
			update = false
		}
	}

	//TODO: need to merge comments and classify censored comments
	// classify if video is now censored at facebook with a flag and never remove comments from our database in that case

	if update {

		_, err = r.Table("sentiment_report").Get(theReport.ID).Update(simpleReport).RunWrite(Session)
		if err != nil {
			log.Println("ERROR: UPDATE sentiment_report TABLE JUST FAILED :|")
		}

		_, err = r.Table("sentiment").Get(theReport.ID).Update(commentsReport).RunWrite(Session)
		if err != nil {
			log.Println("ERROR: UPDATE sentiment TABLE JUST FAILED :|")
		}

	} else {

		log.Println("save first time report into sentiment_report")

		_, err = r.Table("sentiment_report").Insert(simpleReport, r.InsertOpts{Conflict: "update"}).RunWrite(Session)
		if err != nil {
			log.Println("ERROR: INSERT sentiment_report TABLE JUST FAILED :|")
		}

		log.Println("save first time report into sentiment")
		_, err = r.Table("sentiment").Insert(commentsReport, r.InsertOpts{Conflict: "update"}).RunWrite(Session)
		if err != nil {
			log.Println("ERROR: INSERT sentiment TABLE JUST FAILED :|")
		}
	}

	jsonBytes, err = json.Marshal(&theReport)
	if err != nil {
		fmt.Println("Error: ", err)
		w.Write([]byte(`[]`))
		return
	}

	//save to redis
	Client.Set("report_" + postURL, jsonBytes, time.Hour * 1)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

}

func CreateSimpleReport(theReport model.Report) (returnReport model.Report) {

	returnReport.ID = theReport.ID
	returnReport.URL = theReport.URL
	returnReport.Type = theReport.Type
	returnReport.Title = theReport.Title
	returnReport.PublishedAt = theReport.PublishedAt
	returnReport.TotalComments = theReport.TotalComments
	returnReport.CollectedComments = theReport.CollectedComments
	returnReport.CommentCoveragePercent = theReport.CommentCoveragePercent
	returnReport.CommentAvgPerDay = theReport.CommentAvgPerDay
	returnReport.Keywords = theReport.Keywords
	returnReport.Sentiment = theReport.Sentiment
	returnReport.Metadata = theReport.Metadata

	return returnReport
}

func CreateCommentsReport(theReport model.Report) (returnReport model.Report) {

	returnReport.ID = theReport.ID
	returnReport.SentimentList = theReport.SentimentList
	returnReport.TopComments = theReport.TopComments

	return returnReport
}

func SentimentHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.
	lang := req.URL.Query().Get("lang")

	//validate inputs
	if lang == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing lang"})
		w.Write(jsonBytes)
		return
	}

	notTrained := LoadTrainingData(lang)
	if !notTrained {
		log.Println("Unable to train the engine :| No dougnuts for you today o__0")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("lang , ", lang)
	if lang != "pt" && lang != "en" && lang != "fr" && lang != "es" && lang != "it"  {
		errMsg := "Error: SentimentHandler Language " + lang + " not yet supported, use lang={en|pt|es|it|fr} eg lang=en"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	text := req.Form["text"]
	if len(text) == 0 {
		log.Println("Error: SentimentHandler text 404 not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//classify text
	comment := model.Comment{}
	comment.Content = string(text[0])
	comment.Language = lang
	comment.Sentiment = comment.GetSentiment()

	if comment.Sentiment == "" {
		log.Println("Error: SentimentHandler could not get the sentiment for this text :| ", comment.Content)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//marshal comment
	jsonBytes, err := json.Marshal(&comment)
	if err != nil {
		fmt.Println("Error: SentimentHandler Marshal -> ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func AllowOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	//TODO: add origin validation
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
}
