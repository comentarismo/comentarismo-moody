package main

import (
	"comentarismo-moody/server"
	"os"
)

//Host eg: localhost:3001 ; or website.com
var Port = os.Getenv("PORT")

func main() {
	if Port == "" {
		Port = "3000"
	}
	server.StartServer(Port)
}
