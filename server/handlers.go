package server

import (
//util "comentarismo-moody/util"
	"encoding/json"
	"log"
	"net/http"
	"gopkg.in/redis.v3"
	"time"
)

func init() {

}

func MoodyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ReportHandler")
	postURL := r.URL.Query().Get("vid")
	refresh := r.URL.Query().Get("refresh")

	if (refresh == "") {
		//get from redis if available
		//serve from redis
		cached, err := Client.Get("report_" + postURL).Result()
		if err == redis.Nil {
			log.Println("report_" + postURL + " does not exists")
		} else if err != nil {
			panic(err)
		} else {
			log.Println("Serve from cache")
			w.Write([]byte(cached))
			return
		}
	}

	var jsonBytes []byte

	if postURL != "" {
		log.Println(postURL)
		jsonBytes = RunReport(postURL)
		//save to redis
		Client.Set("report_" + postURL, jsonBytes, time.Hour * 24)

	} else {
		jsonBytes, _ = json.Marshal(WebError{Error: "Missing post URL."})
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
