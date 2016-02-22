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

		postURL := "https://www.youtube.com/watch?v=gKI78VOh4Aw"
		server.Training = "../static/training/afinn-111.csv"
		raw := server.RunReport(postURL)
		So(raw,ShouldNotBeNil)
		So(len(raw),ShouldBeGreaterThan,0)
		log.Println(len(raw))
		report := model.Report{}

		err := json.Unmarshal(raw, &report)
		if err != nil {
			log.Println(err)
		}
		So(report.TotalComments, ShouldBeGreaterThan, 0)
		log.Println("Total comments ",report.TotalComments)


		So(len(report.Sentiment), ShouldBeGreaterThan, 0)
		So(len(report.SentimentList), ShouldBeGreaterThan, 0)

		log.Println("Sentiment Graph ")
		for _,s := range report.Sentiment {
			log.Println(s)
		}

		log.Println("Sentiment List ")
		for k,s := range report.SentimentList {
			log.Println(k)
			log.Printf("%v",s)
		}

	})

}
