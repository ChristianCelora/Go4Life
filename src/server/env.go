package server

import (
	"log"
	"path/filepath"
	"sync"
)

const (
	SERVER_PORT = 8090
)

type ServerEnv struct {
	Port             int
	Tempalate_folder string
}

var lock = &sync.Mutex{}
var env *ServerEnv

func GetEnv() *ServerEnv {
	if env == nil {
		lock.Lock()
		defer lock.Unlock()
		if env == nil {
			log.Printf("Init server env")
			env = initServer()
		}
	}
	return env
}

func initServer() *ServerEnv {
	// @refactor: read from a .env file
	template_folder, err := filepath.Abs("./templates/")
	if err != nil {
		log.Fatal("Env Creation failed: Cannot retrieve template folder Abs path")
	}
	return &ServerEnv{
		Port:             SERVER_PORT,
		Tempalate_folder: template_folder,
	}
}
