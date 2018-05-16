package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func verifyToken(resp http.ResponseWriter, req *http.Request) {
	// verify token, TODO: create a package for token generate / verify
	isTokenValid := true

	// TODO: Add route mapping
	if isTokenValid {
		// TODO: redirect for corresponding service
		resp.Write([]byte("sucess " + req.URL.Path))
	} else {
		// return 401
		resp.WriteHeader(401)
		resp.Write([]byte("Invalid Token"))
	}

	// TODO: read about defer
}

func main() {
	argsProgram := os.Args[1:]

	// TODO: Adjust deploy with the config.json for router
	// TODO: add support for multiple versions
	version := argsProgram[0]

	// TODO: Remove mux lib and test docker again
	router := mux.NewRouter()
	router.PathPrefix("/"+version).Methods("GET", "PUT", "DELETE", "POST").HandlerFunc(verifyToken)

	// TODO: define requests timeout
	if err := http.ListenAndServe(":80", router); err != nil {
		panic(err)
	}
}
