package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"../constants"
)

type log struct {
	Message string `json:"message,omitempty"`
}

func LogNow(message string) {
	fmt.Println(message)
	sendLog(message)
}

func Log(message string) {
	fmt.Println(message)
	go sendLog(message)
}

func sendLog(message string) {
	newLog := log{Message: message}
	logJSON, err := json.Marshal(newLog)
	if err != nil {
		fmt.Println(constants.ErrorSenddingLog, err)
	}

	logReq, err := http.NewRequest("POST", constants.HostLog, bytes.NewReader(logJSON))
	if err != nil {
		fmt.Println(constants.ErrorSenddingLog, err)
	}

	client := &http.Client{Timeout: constants.DefaultTimeout}
	_, err = client.Do(logReq)
	if err != nil {
		fmt.Println(constants.ErrorSenddingLog, err)
	}
}
