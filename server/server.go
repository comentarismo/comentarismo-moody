package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	//"strconv"

	r "github.com/dancannon/gorethink"
	"github.com/gorilla/pat"
	//"github.com/dchest/captcha"
	gracehttp "github.com/facebookgo/grace/gracehttp"

	model "comentarismo-moody/model"
	redis "gopkg.in/redis.v3"
	"strings"
)

var (
	router *pat.Router
	//Session defines the rethinkdb session
	Session *r.Session
	//Cache
	Client *redis.Client
)

//DB eg: localhost:28015 : or website.com
var DB = os.Getenv("db")
var DB2 = os.Getenv("db2")

//Table eg: 3001
var Table = os.Getenv("table")

var maxopen = os.Getenv("maxopen")
var maxidle = os.Getenv("maxidle")
var SEARCH = os.Getenv("search")
var RESIZER = os.Getenv("imgresizer")

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
	if DB == "" {
		DB = "192.168.0.42:28015"
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
	if SEARCH == "" {
		SEARCH = "localhost:8099"
	}
	if RESIZER == "" {
		RESIZER = "localhost:3666"
	}
	if REDIS_HOST == "" {
		REDIS_HOST = "192.168.0.42:6379"
	}
	if REDIS_PASS == "" {
	}
	if Training == "" {
		Training = "./static/training/afinn-111.csv"
	}

	//maxo, _ := strconv.Atoi(maxopen)
	//maxi, _ := strconv.Atoi(maxidle)

	//if DB2 != "" {
	//	log.Println("Loading cluster "+DB+" "+DB2)
	//	Session, err = r.Connect(r.ConnectOpts{
	//		Addresses: []string{DB, DB2},
	//		Database: Table,
	//		MaxOpen:  maxo,
	//		MaxIdle: maxi,
	//		DiscoverHosts: false,
	//	})
	//} else {
	//	log.Println("Loading standalone "+DB)
	//	Session, err = r.Connect(r.ConnectOpts{
	//		Address: DB,
	//		Database: Table,
	//		MaxOpen:  maxo,
	//		MaxIdle: maxi,
	//	})
	//
	//}
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}

	//captcha.SetCustomStore(NewMemoryStore())

	//Client = redis.NewClient(&redis.Options{
	//	Addr:     REDIS_HOST,
	//	Password: REDIS_PASS, // no password set
	//	DB:       0,  // use default DB
	//})

	//pong, err := Client.Ping().Result()
	//log.Println(pong, err)

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

	//r.Handle("/captcha/{id}.{ext}", CaptchaServer(120, 80))

	//r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/api", ReportHandler)

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

	ss := http.StripPrefix("", http.FileServer(http.Dir("./templates/")))
	r.PathPrefix("/").Handler(ss)

	log.Println("Listening on :" + Host + "...")

	return r
}
