package main

import (
	"net/http"
	"os"

	"./auth/"
	"./discovery/"

	"github.com/gorilla/mux"
)

func redirect(resp http.ResponseWriter, req *http.Request) {

	reqURL := discovery.GetServiceURL(req)
	if reqURL != "" {
		resp.WriteHeader(404)
		resp.Write([]byte("Invalid path"))
		return
	}

	if discovery.IsRoutePublic(req.URL.Path) || auth.IsTokenValid(req.Header.Get("token")) {

		redirectResp, err := makeClientRequests(req, reqURL)

		if err != nil {
			resp.WriteHeader(500)
			return
		}

		var bytes []byte
		redirectResp.Body.Read(bytes)

		resp.WriteHeader(redirectResp.StatusCode)
		resp.Write(bytes)
	} else {
		resp.WriteHeader(401)
		resp.Write([]byte("Invalid Token"))
	}
}

// TODO: Verify how work auth routes
func loginRequest(resp http.ResponseWriter, req *http.Request) {
	url := discovery.GetServiceURL(req)
	redirectResp, err := makeClientRequests(req, url)

	if err == nil && redirectResp.StatusCode == 200 {
		resp.WriteHeader(200)
		resp.Write([]byte(auth.GenerateToken(redirectResp.Header.Get("id"))))
	} else {
		resp.WriteHeader(500)
	}
}

// TODO: dont create a new client for every request
func makeClientRequests(req *http.Request, url string) (*http.Response, error) {
	redirectReq, _ := http.NewRequest(req.Method, url, req.Body)
	client := &http.Client{}
	return client.Do(redirectReq)
}

// TODO: check all error
// TODO: make a method log
func main() {
	argsProgram := os.Args[1:]

	port := ":80"
	if len(argsProgram) > 0 {
		port = argsProgram[0]
	}

	router := mux.NewRouter()
	router.HandleFunc("/auth", loginRequest)

	// TODO: set this method to be used async
	discovery.StartDiscoveryService(router, redirect)

	// TODO: define requests timeout
	if err := http.ListenAndServe(port, router); err != nil {
		panic(err)
	}
}
