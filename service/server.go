package main

import (
	"net/http"
	"os"

	"./auth"
	"./constants"
	"./db"
	"./log"

	"github.com/gorilla/mux"
)

func main() {
	argsProgram := os.Args[1:]

	port := constants.DefaultServerPort
	if len(argsProgram) > 0 {
		port = argsProgram[0]
	}

	router := mux.NewRouter()
	router.HandleFunc(constants.PathLogin, loginRequest).Methods("POST")
	router.HandleFunc(constants.PathRegister, registerRequest).Methods("POST")
	router.PathPrefix(constants.PathAPI).HandlerFunc(redirect)

	db.Start()
	auth.CreateSecret()

	log.Log(constants.MessageStartingServer + " - " + port)

	if err := http.ListenAndServe(port, router); err != nil {
		panic(err)
	}
}
