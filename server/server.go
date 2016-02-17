package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/pat"
	gracehttp "github.com/facebookgo/grace/gracehttp"

	model "comentarismo-moody/model"
	"strings"
	"comentarismo-moody/providers/facebook"
	"comentarismo-moody/providers/youtube"
)

var (
	router *pat.Router
)


var REDIS_HOST = os.Getenv("redis_host")
var REDIS_PASS = os.Getenv("redis_pass")

var Host = os.Getenv("host")

//Port eg: 3000
var Port = os.Getenv("PORT")

var Training = os.Getenv("training")

func init() {
	//var err error

	if Port == "" {
		Port = "3000"
	}
	if Host == "" {
		Host = "http://localhost" + ":" + Port
	}
	if REDIS_HOST == "" {
		REDIS_HOST = "192.168.0.42:6379"
	}
	if REDIS_PASS == "" {
	}
	if Training == "" {
		Training = "./static/training/afinn-111.csv"
	}

	trainingFiles := strings.Split(Training, ",")
	for _, path := range trainingFiles {
		model.LoadTrainingData(path)
	}
}

//NewServer return pointer to new created server object
func NewServer(Port string) *http.Server {
	router = InitRouting()
	return &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}
}

//StartServer start and listen @server
func StartServer(Port string) {
	log.Println("Starting server")
	s := NewServer(Port)
	fmt.Println("Server starting --> " + Port)

	err := gracehttp.Serve(
		s,
	)

	if err != nil {
		log.Fatalln("Error: %v", err)
		os.Exit(0)
	}
}

func InitRouting() *pat.Router {
	r := pat.New()

	InitProviders();

	r.HandleFunc("/moody", MoodyHandler)

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

	ss := http.StripPrefix("", http.FileServer(http.Dir("./templates/")))
	r.PathPrefix("/").Handler(ss)

	log.Println("Listening on :" + Host + "...")

	return r
}

func InitProviders(){
	model.UseProviders(
		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), Host + "/auth/facebook/callback"),
		youtube.New(os.Getenv("YOUTUBE_KEY"), os.Getenv("YOUTUBE_SECRET"), Host + "/auth/lastfm/callback"),
	)
}
