package main

import (
	"golife/api"
	"golife/server"
	"log"
	"log/slog"
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
	http.HandleFunc("/api/healtcheck", api.Healtcheck)
}

func main() {
	env := server.GetEnv()
	registerStaticPages()
	registerApiRoutes()
	slog.Info("Listening on", "port", strconv.Itoa(env.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(env.Port), nil)
	if err != nil {
		slog.Error("", "msg", err)
		log.Fatal(err)
	}
}
