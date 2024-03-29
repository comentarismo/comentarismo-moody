package server

import (
	//util "comentarismo-moody/util"
	"encoding/json"
	r "github.com/dancannon/gorethink"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"time"

	"comentarismo-moody/lang"
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
		log.Println("MoodyHandler, request will default to lang -> ", lang)
	} else {
		log.Println("MoodyHandler, request will use lang -> ", lang)
	}
	//validate inputs
	if postURL == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing post URL"})
		w.Write(jsonBytes)
		return
	}

	log.Println("MoodyHandler, ", postURL, refresh)
	var jsonBytes []byte

	UUID, err := UUID(postURL)
	if err != nil {
		log.Println("MoodyHandler, Could not SetID. report_"+postURL+" UUID failed ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if refresh == "" || refresh != "true" {
		//get from redis if available
		//serve from redis
		valid := true

		cached, err := Client.Get("report_" + postURL).Result()
		if err == redis.Nil {
			log.Println("MoodyHandler, report_" + postURL + " does not exists in cache")
		} else if err != nil {
			log.Println("Error: REDIS Error, MoodyHandler, report_"+postURL+" does not exists in cache -> ", err)
		} else {
			log.Println("MoodyHandler, Serve from cache -> ", postURL)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cached))
			return
		}

		log.Println("MoodyHandler, not in redis, is it on rethinkdb ??")

		/** BEGIN Code to unmarshal the url into a valid ID that we can use in our internal database **/
		domain, urlParts := parseURL(postURL)
		if domain == "" || len(urlParts) == 0 {
			log.Println("Error: MoodyHandler, parseURL -> report_" + postURL + " does not exists ")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		provider, err := model.GetProvider(domain)
		if err != nil {
			log.Println("MoodyHandler, Could not GetProvider. report_"+postURL+" does not exists ", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = provider.SetID(urlParts)
		if err != nil {
			log.Println("MoodyHandler, Could not SetID. report_"+postURL+" does not exists ", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		theReport := model.Report{}
		provider.SetReport(&theReport, model.CommentList{})
		/** END Code to unmarshal the url into a valid ID that we can use in our internal database **/

		simpleReport := model.YoutubeReport{}

		commentsReport := model.Report{}

		log.Println("MoodyHandler, r.Table(sentiment_report).Get(" + UUID + ")")
		/** BEGIN recover the simple sentiment report **/
		res, err := r.Table("sentiment_report").Get(UUID).Run(Session)

		//if found rethinkdb sentiment_report, query all comments and save cache redis
		if err != nil || res == nil {
			log.Println("MoodyHandler, Could not get report_"+postURL+" does not exists rethinkdb ", err)
			valid = false
		} else if res.IsNil() {
			log.Println("MoodyHandler, Could not SetID. report_"+postURL+" does not exists rethinkdb ", err)
			valid = false
		} else {
			err = res.One(&simpleReport)
			if err != nil {
				log.Println("Error: MoodyHandler, res.One(&simpleReport), ", err)
			}
			if simpleReport.ID == "" || simpleReport.TotalComments == 0 {
				log.Println("MoodyHandler, ", simpleReport)
				valid = false
				log.Println("MoodyHandler, Could NOT GET ID. report_" + postURL + " does not exists rethinkdb")
			} else {
				log.Println("MoodyHandler, OK GET ID. report_" + postURL + " exists rethinkdb ")

			}
			defer res.Close()
		}
		/** END recover the simple sentiment report **/

		/** BEGIN recover the comments sentiment report **/
		if valid {
			log.Println("MoodyHandler, load sentiment_report ok, going to load all comments now, this may break :| ")
			log.Println("MoodyHandler, r.Table(sentiment).Get(" + UUID + ").Pluck(sentimentlist, topcomments)")
			res2, err := r.Table("sentiment").Get(UUID).Pluck("sentimentlist", "topcomments").Run(Session)
			if err != nil || res2 == nil {
				log.Println("MoodyHandler, err != nil || res2 == nil, sentimentlist, Could not get report_"+postURL+" does not exists rethinkdb ", err)
				valid = false
			} else if res2.IsNil() {
				log.Println("Error: res2.IsNil(), MoodyHandler, sentimentlist, Could not SetID. report_"+postURL+" does not exists rethinkdb ", err)
				valid = false
			} else {
				defer res2.Close()
				res2.One(&commentsReport)
			}
		}
		/** END recover the comments sentiment report **/

		if valid {

			log.Println("MoodyHandler, SentimentList: ", len(commentsReport.SentimentList))
			//log.Println("SampleComments: ",len(item.SampleComments))
			log.Println("MoodyHandler, TopComments: ", len(commentsReport.TopComments))

			//validate there is a sentiment list and top comments is filled
			if len(commentsReport.SentimentList) > 0 && len(commentsReport.TopComments) > 0 {

				//merge both into one

				simpleReport.SentimentList = commentsReport.SentimentList
				simpleReport.TopComments = commentsReport.TopComments

				jsonBytes, err = json.Marshal(&simpleReport)

				if err != nil {
					log.Println("Error: MoodyHandler, Marshal, OMG we could not get this SentimentList to work :| No dougnuts for you today o__0 ", err)
				} else {

					log.Println("MoodyHandler, loading from rethinkdb went ok :D ")

					Client.Set("report_"+postURL, jsonBytes, time.Hour*1)
					log.Println("MoodyHandler, save to redis ok")

					w.Header().Set("Content-Type", "application/json")
					w.Write(jsonBytes)
					return
				}
			} else {
				log.Println("Error: MoodyHandler, len(commentsReport.SentimentList) > 0 && len(commentsReport.TopComments) > 0, OMG we could not get this SentimentList to work :| No dougnuts for you today o__0 ")
			}
		} else {
			log.Println("Error: MoodyHandler, !valid, OMG we could not get this validation to work :| No dougnuts for you today o__0 ")
		}

	}

	log.Println("MoodyHandler, REFRESH -> Will generate a new report -->", postURL)

	quickReport, err := RunQuickReport(postURL, lang)
	if err != nil {
		log.Println("Error: MoodyHandler, RunQuickReport() ", err)
		w.WriteHeader(http.StatusInternalServerError)
		jsonBytes, _ := json.Marshal(WebError{Error: "RunQuickReport failed :| -> "})
		w.Write(jsonBytes)
		return
	}

	quickReport.ID = UUID

	go func() {
		theReport, err := RunReport(postURL, lang)
		if err != nil {
			log.Println("Error: MoodyHandler, RunReport() ", err)
			//w.WriteHeader(http.StatusInternalServerError)
			//jsonBytes, _ := json.Marshal(WebError{Error: "RunReport failed :| -> "})
			//w.Write(jsonBytes)
			return
		}
		//save to rethinkdb
		err = SaveReport(theReport, postURL, UUID)
		if err != nil {
			log.Println("Error: MoodyHandler, SaveReport() ", err)
			return
		}
	}()

	jsonBytes, err = json.Marshal(&quickReport)
	if err != nil {
		log.Println("Error: MoodyHandler, quickReport, ", err)
		w.Write([]byte(`[]`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

}

func CreateSimpleReport(theReport model.Report) (returnReport model.Report) {

	returnReport.ID = theReport.ID
	returnReport.URL = theReport.URL
	returnReport.Type = theReport.Type
	returnReport.Title = theReport.Title
	returnReport.Date = theReport.Date
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
	req.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.
	lang := req.URL.Query().Get("lang")

	//validate inputs
	if lang == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing lang"})
		w.Write(jsonBytes)
		return
	}

	//log.Println("lang , ", lang)
	if lang != "pt" && lang != "en" && lang != "fr" && lang != "es" && lang != "it" && lang != "hr" && lang != "ru" && lang != "nl" && lang != "cn" {
		errMsg := "Error: SentimentHandler Language " + lang + " not yet supported, use lang={en|pt|es|it|fr|hr|ru|nl} eg lang=en"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	notTrained := LoadTrainingData(lang)
	if !notTrained {
		log.Println("Unable to train the engine :| No dougnuts for you today o__0")
		w.WriteHeader(http.StatusInternalServerError)
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
	comment.Comment = string(text[0])
	comment.Language = lang
	comment.Sentiment = comment.GetSentiment()

	if comment.Sentiment == 0 {
		log.Println("Error: SentimentHandler could not get the sentiment for this text :| ", comment.Comment)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//  field comment removed to save some bytes on the network
	comment.Comment = string(text[0])

	//marshal comment
	jsonBytes, err := json.Marshal(&comment)
	if err != nil {
		log.Println("Error: SentimentHandler Marshal -> ", err)
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

func LanguageHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	text := req.Form["text"]
	if len(text) == 0 {
		log.Println("Error: LanguageHandler text 404 not found", req.Form)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//classify language for text
	var returnLang lang.DetectLanguage
	detectedLang, err := lang.Guess(text[0])
	if err != nil {
		returnLang.Error = err.Error()
		log.Println("Error: LanguageHandler could detect lang for this text :| ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		returnLang.Language = detectedLang
		returnLang.Date = time.Now()
		returnLang.Text = text[0]
		log.Println(returnLang)
	}

	//marshal comment
	jsonBytes, err := json.Marshal(&returnLang)
	if err != nil {
		log.Println("Error: LanguageHandler Marshal -> ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
