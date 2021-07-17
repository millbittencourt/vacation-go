package main

import (
	"log"
	"net/http"

	"mille.com/todo/data"
	"mille.com/todo/route"
)

func main() {
	data.InitDatabase()
	log.Print("server is on mami")
	http.ListenAndServe(":8888", route.RegisterRoute())
}
