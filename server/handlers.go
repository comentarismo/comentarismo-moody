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
	log.Println("ReportHandler ",postURL,refresh)
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

		domain, urlParts := parseURL(postURL)
		if domain == "" || len(urlParts) == 0 {
			log.Println("report_" + postURL + " does not exists")
			log.Println("404 not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		provider, err := model.GetProvider(domain)
		if err != nil {
			log.Println("Could not GetProvider. report_" + postURL + " does not exists",err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = provider.SetID(urlParts)
		if err != nil {
			log.Println("Could not SetID. report_" + postURL + " does not exists",err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		theReport := model.Report{}
		provider.SetReport(&theReport, model.CommentList{})

		//if found rethinkdb save cache redis
		res, err := r.Table("sentiment").Get(theReport.ID).Run(Session)
		if err != nil {
			log.Println("Could not get report_" + postURL + " does not exists",err)
			valid = false
		}
		if res.IsNil() {
			log.Println("Could not SetID. report_" + postURL + " does not exists",err)
			valid = false
		}
		defer res.Close()

		if valid {
			item := model.Report{}
			res.One(&item)
			if (item.ID == "") {
				valid = false
				log.Println("Could GET ID. report_" + postURL + " does not exists", err)
			}

			log.Println("SentimentList: ",len(item.SentimentList))
			//log.Println("SampleComments: ",len(item.SampleComments))
			log.Println("TopComments: ",len(item.TopComments))

			if len(item.SentimentList) == 0 {
				log.Println("RETHINKDB COULD NOT GET DATA CORRECTLY!!! karaidiasa!!!!!")
			}

			if len(item.SentimentList) > 0 {
				jsonBytes, err = json.Marshal(item)

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
			}
		}

	}


	if postURL != "" {
		log.Println(postURL)
		theReport, err := RunReport(postURL)
		if err != nil {
			fmt.Println("Error: ",err)
		}
		//save to rethinkdb

		update := true

		res, err := r.Table("sentiment").Get(theReport.ID).Run(Session)
		if err != nil {
			update = false
		}
		if res.IsNil() {
			update = false
		}
		defer res.Close()

		item := model.Report{}
		res.One(&item)
		if(item.ID == ""){
			update = false
		}

		if update {
			log.Println("lets update sentiment report")

			//TODO: need to merge comments and classify censored comments as well

			_, err = r.Table("sentiment").Get(theReport.ID).Update(theReport).RunWrite(Session)
			if err != nil {
				log.Println("ERROR: UPDATE SENTIMENT TABLE JUST FAILED :|")
			}
		}else {
			log.Println("save first time sentiment report")
			_, err = r.Table("sentiment").Insert(theReport).RunWrite(Session)
			if err != nil {
				log.Println("ERROR: INSERT SENTIMENT TABLE JUST FAILED :|")
			}
		}

		jsonBytes, err = json.Marshal(theReport)
		if err != nil {
			fmt.Println("Error: ",err)
			w.Write([]byte(`[]`))
			return
		}

		//save to redis
		Client.Set("report_" + postURL, jsonBytes, time.Hour * 1)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	} else {
		jsonBytes, _ = json.Marshal(WebError{Error: "Missing post URL."})
		w.Write(jsonBytes)
		return
	}

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
