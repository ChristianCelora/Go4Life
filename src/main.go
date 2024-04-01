package main

import (
	"github.com/tbxark/g4vercel"
	"golife/api"
	"log"
	"net/http"
	"strconv"
)

const (
	SERVER_PORT = 8090
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
	print("Go 4 Life")

	registerStaticPages()
	registerApiRoutes()
	log.Print("Listening on port " + strconv.Itoa(SERVER_PORT))
	err := http.ListenAndServe(":"+strconv.Itoa(SERVER_PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Vercel is serverless
func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/render", api.RenderMatrix)

	server.Handle(w, r)
}
