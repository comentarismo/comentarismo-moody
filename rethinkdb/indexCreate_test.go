package rethinkdb

import (
	. "github.com/smartystreets/goconvey/convey"
	r "gopkg.in/gorethink/gorethink.v3"

	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

var session *r.Session
var RETHINKDB_HOST = os.Getenv("RETHINKDB_HOST")
var RETHINKDB_PORT = os.Getenv("RETHINKDB_PORT")
var Table = os.Getenv("table")
var DB = ""

//RETHINKDB_HOST=g7-box go test -v rethinkdb/indexCreate_test.go
//RETHINKDB_HOST=g7-box table=test go test -v rethinkdb/indexCreate_test.go

func init() {
	if RETHINKDB_HOST == "" {
		RETHINKDB_HOST = "g7-box"
	}

	if RETHINKDB_PORT == "" {
		RETHINKDB_PORT = "28015"
	}

	if Table == "" {
		Table = "comentarismo"
	}

	DB = RETHINKDB_HOST + ":" + RETHINKDB_PORT

}

//TABLE sentiment_report
//collectedcomments
//commentavgperday
//commentcoveragepercent
//date
//id
//keywords
//metadata.channelid
//metadata.channeltitle
//metadata.clientkey
//metadata.id
//metadata.language
//metadata.publishedat
//metadata.thumbnail
//metadata.title
//metadata.totalcomments
//metadata.updatedAt
//metadata.videoviews
//publishedat
//sentiment
//title
//totalcomments
//type
//updatedAt
//url

//TABLE sentiment
//sentimentlist.Awesome!
//sentimentlist.Bad
//sentimentlist.Eh
//sentimentlist.Good
//sentimentlist.Like It
//sentimentlist.Loved It
//sentimentlist.Neutral
//sentimentlist.Not Good
//sentimentlist.OK
//sentimentlist.Sucks
//sentimentlist.Terrible!
//sentimentlist.Unknown
//topcomments
//updatedAt

//test will create all the tables and indexes on the rethinkdb instance target
//and assert that all indexes have been created after it
func TestCreateIndexesMoody(t *testing.T) {

	var err error
	Convey("Given I create all databases for comentarismo-moody", t, func() {

		_, err = r.DBCreate(Table).RunWrite(session)
		if err != nil {
			So(err.Error(), ShouldResemble, "gorethink: Database `"+Table+"` already exists. in:\nr.DBCreate(\""+Table+"\")")
		}

		Convey("When I create all tables for comentarismo-moody", func() {

			r.DB(Table).TableCreate("sentiment").Exec(session)
			r.DB(Table).TableCreate("sentiment_report").Exec(session)
			r.DB(Table).TableCreate("commentaries_sentiment_report").Exec(session)

			Convey("Then I create all indexes for comentarismo-moody", func() {

				/**
				 * sentiment_report & sentiment
				 */
				r.DB(Table).Table("sentiment_report").IndexCreate("type").Exec(session)
				r.DB(Table).Table("sentiment_report").IndexCreate("publishedat").Exec(session)
				r.DB(Table).Table("sentiment_report").IndexCreate("date").Exec(session)

				r.DB(Table).Table("sentiment_report").IndexCreateFunc(
					"operator_date",
					[]interface{}{r.Row.Field("operator"), r.Row.Field("date")},
				).Exec(session)

				//r.DB(Table).Table("sentiment_report").IndexWait().Exec(session)
				//r.DB(Table).Table("commentator_sentiment_report").IndexWait().Exec(session)

			})
		})

	})

}

func TestMain(m *testing.M) {
	// seed randomness for use with tests
	rand.Seed(time.Now().UTC().UnixNano())

	testSetup(m)
	res := m.Run()
	testTeardown(m)

	os.Exit(res)
}

func testSetup(m *testing.M) {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  DB,
		Database: Table,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func testTeardown(m *testing.M) {
	session.Close()
}
