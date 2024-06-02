package main

import (
	"golife/api"
	"golife/server"
	"log"
	"net/http"
	"strconv"
)

func registerStaticPages() {
	fs := http.FileServer(http.Dir("./static/view"))
	http.Handle("/", fs)
}

func registerApiRoutes() {
	http.HandleFunc("/api/render", api.RenderMatrix)
	http.HandleFunc("/api/step", api.GetNextStep)
}

func main() {
	env := server.GetEnv()
	registerStaticPages()
	registerApiRoutes()
	log.Printf("Listening on port %s", strconv.Itoa(env.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(env.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
