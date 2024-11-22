package main

import (
	"Hausuebung-I/repo/mongodb"
	"Hausuebung-I/transport/rest"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nilspolek/goLog"
)

func main() {
	repo := mongodb.New()
	router := mux.NewRouter()

	rest := rest.New(router, &repo)
	if err := http.ListenAndServe(":8080", rest.Router); err != nil {
		goLog.Error("%s", err.Error())
	}
}
