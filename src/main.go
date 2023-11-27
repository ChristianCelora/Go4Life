package main

import (
	"net/http"
	"golife/api"
	"strconv"
	"log"
)

const (
	SERVER_PORT = 8090
)

func registerStaticPages() {
	fs := http.FileServer(http.Dir("./view"))
	http.Handle("/", fs)
}

func registerApiRoutes() {
	http.HandleFunc("/render", api.RenderMatrix)
}

func main() {
	print("Go 4 Life")

	registerStaticPages()
	registerApiRoutes()
	log.Print("Listening on port " + strconv.Itoa(SERVER_PORT))
	err := http.ListenAndServe(":" + strconv.Itoa(SERVER_PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}