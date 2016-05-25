package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/pat"
	"github.com/facebookgo/grace/gracehttp"
	redis "gopkg.in/redis.v3"
	r "github.com/dancannon/gorethink"

	"comentarismo-moody/providers/facebook"
	"comentarismo-moody/providers/youtube"
	"comentarismo-moody/providers/instagram"
	"comentarismo-moody/providers/vinevideo"

	"comentarismo-moody/model"
	"strconv"
)

var (
	Session *r.Session
	router *pat.Router
	Client *redis.Client
)

var REDIS_HOST = os.Getenv("REDIS_HOST")
var REDIS_PORT = os.Getenv("REDIS_PORT")
var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

var Host = os.Getenv("host")

//Port eg: 3000
var Port = os.Getenv("PORT")
var DB = os.Getenv("db")

//Table eg: 3001
var Table = os.Getenv("table")

var maxopen = os.Getenv("maxopen")
var maxidle = os.Getenv("maxidle")

func init() {
	//var err error

	if DB == "" {
		DB = "g7-box:28015"
	}

	if Table == "" {
		Table = "test"
	}
	if maxopen == "" {
		maxopen = "80"
	}
	if maxidle == "" {
		maxidle = "40"
	}
	maxo, _ := strconv.Atoi(maxopen)
	maxi, _ := strconv.Atoi(maxidle)

	log.Println("Loading standalone " + DB)
	var err error
	Session, err = r.Connect(r.ConnectOpts{
		Address: DB,
		Database: Table,
		MaxOpen:  maxo,
		MaxIdle: maxi,
	})
	if err != nil {
		log.Fatalln("Error: %v", err)
		os.Exit(0)
	}

	if Port == "" {
		Port = "3000"
	}
	if Host == "" {
		Host = "http://localhost" + ":" + Port
	}
	if REDIS_HOST == "" {
		REDIS_HOST = "g7-box"
	}
	if REDIS_PORT == "" {
		REDIS_HOST = "6379"
	}
	if REDIS_PASSWORD == "" {
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Password: REDIS_PASSWORD, // no password set
		DB:       0, // use default DB
	})

	pong, err := Client.Ping().Result()
	if err != nil {
		log.Fatalln("Error: %v", err)
		os.Exit(0)
	}
	log.Println(pong, err)
}

//NewServer return pointer to new created server object
func NewServer(Port string) *http.Server {
	router = InitRouting()
	return &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}
}

func initProviders() {
	log.Println("Going to register providers")
	// Provider is the interface for all the various 3rd party Provider types (YouTube, Facebook, etc...)
	model.UseProviders(
		facebook.New(FACEBOOK_KEY, FACEBOOK_SECRET, Host + "/auth/facebook/callback"),
		youtube.New(YOUTUBE_KEY, YOUTUBE_SECRET, Host + "/auth/youtube/callback"),
		instagram.New(INSTAGRAM_KEY, INSTAGRAM_SECRET, Host + "/auth/instagram/callback"),
		vinevideo.New(VINEVIDEO_KEY, VINEVIDEO_SECRET, Host + "/auth/vinevideo/callback"),
	)
}
//StartServer start and listen @server
func StartServer(Port string) {
	log.Println("Starting server")
	s := NewServer(Port)
	initProviders()
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

	r.HandleFunc("/moody", MoodyHandler)

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

	ss := http.StripPrefix("", http.FileServer(http.Dir("./templates/")))
	r.PathPrefix("/").Handler(ss)

	log.Println("Listening on :" + Host + "...")

	return r
}