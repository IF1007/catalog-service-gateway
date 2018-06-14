package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"./auth"
	"./constants"
	"./db"
	"./dns"
	"./log"
)

func redirect(resp http.ResponseWriter, req *http.Request) {

	reqURL := dns.GetServiceURL(req.URL.Path)
	if reqURL == "" {
		resp.WriteHeader(404)
		return
	}

	tokenEncrypt := req.Header.Get(constants.AttrToken)
	id := auth.GetIdFromToken(tokenEncrypt)
	if !dns.IsRoutePublic(req.URL.Path) || id != "" {

		log.Log(id + constants.IsRequestingTo + reqURL)
		req.Header.Set("ID", id)
		redirectResp, err := makeClientRequests(req, reqURL)

		if err != nil {
			log.Log(constants.ErrorRequesting + err.Error())
			resp.WriteHeader(500)
			return
		}

		bytes, _ := ioutil.ReadAll(redirectResp.Body)
		if redirectResp.StatusCode == 500 {
			resp.WriteHeader(500)
			log.Log(constants.ErrorCrud + string(bytes))
		} else {
			resp.WriteHeader(redirectResp.StatusCode)
			resp.Write(bytes)
		}
	} else {
		resp.WriteHeader(401)
		resp.Write([]byte(constants.ErrorInvalidToken))
	}
}

func loginRequest(resp http.ResponseWriter, req *http.Request) {

	var usrReq db.UserAuth
	_ = json.NewDecoder(req.Body).Decode(&usrReq)

	usr, err := db.GetUserByLoginPass(usrReq.Login, usrReq.Password)
	if err != nil {
		log.Log(constants.ErrorLogin + " - " + err.Error())
		resp.WriteHeader(500)
		return
	}
	if usr != nil {
		log.Log(constants.MessageNewUserLogin + usr.Login)
		resp.WriteHeader(200)
		resp.Write([]byte(auth.GenerateToken(usr.ID)))
	} else {
		resp.WriteHeader(400)
		resp.Write([]byte(constants.ErrorInvalidUserOrPass))
	}
}

func registerRequest(resp http.ResponseWriter, req *http.Request) {

	var usrReq *db.UserAuth

	_ = json.NewDecoder(req.Body).Decode(&usrReq)
	if usrReq == nil {
		resp.WriteHeader(400)
		return
	}

	hasLogin, err := db.HasLogin(usrReq.Login)
	if err != nil {
		log.Log(constants.ErrorRegisterNewUser + " - " + err.Error())
		resp.WriteHeader(500)
		return
	}

	if hasLogin {
		resp.WriteHeader(400)
		resp.Write([]byte(constants.ErrorLoginAlreadyExists))
		return
	}

	invalidParams, err := db.InsertUser(usrReq)

	if invalidParams {
		resp.WriteHeader(400)
	}

	if err != nil {
		log.Log(constants.ErrorRegisterNewUser + " - " + err.Error())
		resp.WriteHeader(500)
	} else {
		log.Log(constants.MessageNewUserCreate + usrReq.Login)
		resp.WriteHeader(200)
	}
}

// TODO: dont create a new client for every request
func makeClientRequests(req *http.Request, url string) (*http.Response, error) {
	redirectReq, _ := http.NewRequest(req.Method, url, req.Body)
	// TODO: add some day a timeout
	client := &http.Client{}
	return client.Do(redirectReq)
}
