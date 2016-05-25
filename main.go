package main

import (
	"comentarismo-moody/server"
	"os"
)

//Host eg: localhost:3003
var Port = os.Getenv("PORT")

func main() {
	if Port == "" {
		Port = "3003"
	}
	server.StartServer(Port)
}
