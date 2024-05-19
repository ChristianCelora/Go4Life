package main

import (
	"golife/api"
	"golife/server"
	"log"
	"net/http"
	"strconv"
)

func registerStaticPages() {
	// views
	fs := http.FileServer(http.Dir("./static/view"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// node modules
	fs = http.FileServer(http.Dir("./node_modules/"))
	http.Handle("/static/node_modules/", http.StripPrefix("/node_modules/", fs))
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
