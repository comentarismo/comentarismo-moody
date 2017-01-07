package test

import (
	"comentarismo-moody/model"
	"comentarismo-moody/providers/youtube"
	"comentarismo-moody/server"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func Test_Youtube_Sentiment(t *testing.T) {

	store := model.InitSentiment("en")
	err := store.Reset()

	Convey("Should RunReport for youtube", t, func() {
		So(err, ShouldBeNil)

		YOUTUBE_KEY := "AIzaSyBameiyxxJw0W27lydpPuPRocfvGza9gXM"

		y := youtube.New(YOUTUBE_KEY, "", "")
		model.UseProviders(y)

		postURL := "https://www.youtube.com/watch?v=hAu5_m26Tsc"
		//server.Training = "../static/training/afinn-111-en.csv"
		theReport, err := server.RunReport(postURL, "en")
		raw, err := json.Marshal(&theReport)

		So(err, ShouldBeNil)
		So(raw, ShouldNotBeNil)
		log.Println("len(raw) ", len(raw))
		So(len(raw), ShouldBeGreaterThan, 0)

		//report := model.Report{}

		//err = json.Unmarshal(raw, &report)
		//So(err,ShouldBeNil)

		log.Println("Total comments ", theReport.TotalComments)
		So(theReport.TotalComments, ShouldBeGreaterThan, 0)

		So(len(theReport.Sentiment), ShouldBeGreaterThan, 0)
		So(len(theReport.SentimentList), ShouldBeGreaterThan, 0)

		log.Println("Sentiment Graph ")
		for _, s := range theReport.Sentiment {
			log.Println(s)
		}

		log.Println("Sentiment List ")
		for k, s := range theReport.SentimentList {
			log.Println(k)
			log.Printf("%v", s)
		}

	})

}
