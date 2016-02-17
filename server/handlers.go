package server

import (
	//util "comentarismo-moody/util"
	"encoding/json"
	"log"
	"net/http"
)

func MoodyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ReportHandler")
	postURL := r.URL.Query().Get("vid")

	var jsonBytes []byte

	if postURL != "" {
		log.Println(postURL)
		jsonBytes = RunReport(postURL)
	} else {
		jsonBytes, _ = json.Marshal(WebError{Error: "Missing post URL."})
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
