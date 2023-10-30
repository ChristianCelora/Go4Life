package main

import (
	"net/http"
	"golife/api"
)

const (
	SERVER_PORT = 8090
)

func registerApiRoutes() {
	http.HandleFunc("/render", api.RenderMatrix)
}

func main() {
	print("Go 4 Life")

	registerApiRoutes()
	http.ListenAndServe(":" + string(SERVER_PORT), nil)
}