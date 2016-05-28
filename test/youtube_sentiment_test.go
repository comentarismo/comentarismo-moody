package test

import (
	"testing"
	"comentarismo-moody/server"
	"comentarismo-moody/model"
	"encoding/json"
	"log"
	"comentarismo-moody/providers/youtube"
	. "github.com/smartystreets/goconvey/convey"

)

func Test_Youtube_Sentiment(t *testing.T) {

	Convey("Should RunReport for youtube", t, func() {

		YOUTUBE_KEY := "AIzaSyBameiyxxJw0W27lydpPuPRocfvGza9gXM"

		y := youtube.New(YOUTUBE_KEY, "", "")
		model.UseProviders(y)

		postURL := "https://www.youtube.com/watch?v=hAu5_m26Tsc"
		server.Training = "../static/training/afinn-111-en.csv"
		theReport, err := server.RunReport(postURL)
		raw, err := json.Marshal(&theReport)

		So(err,ShouldBeNil)
		So(raw,ShouldNotBeNil)
		log.Println("len(raw) ",len(raw))
		So(len(raw),ShouldBeGreaterThan,0)

		//report := model.Report{}

		//err = json.Unmarshal(raw, &report)
		//So(err,ShouldBeNil)

		log.Println("Total comments ",theReport.TotalComments)
		So(theReport.TotalComments, ShouldBeGreaterThan, 0)

		So(len(theReport.Sentiment), ShouldBeGreaterThan, 0)
		So(len(theReport.SentimentList), ShouldBeGreaterThan, 0)

		log.Println("Sentiment Graph ")
		for _,s := range theReport.Sentiment {
			log.Println(s)
		}

		log.Println("Sentiment List ")
		for k,s := range theReport.SentimentList {
			log.Println(k)
			log.Printf("%v",s)
		}

	})

}
