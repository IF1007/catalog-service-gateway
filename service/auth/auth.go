package auth

import (
	"encoding/json"
	"math/rand"
)

type token struct {
	ID   string `json:"id,omitempty"`
	Salt int    `json:"salt,omitempty"`
}

const maxInt = int((^uint(0)) >> 1)

func GenerateToken(id string) string {
	newToken := &token{ID: id, Salt: rand.Intn(maxInt)}

	jsonToken, _ := json.Marshal(newToken)

	// TODO: crypt token
	return string(jsonToken)
}

func GetToken(requestTokenStr string) string {
	// TODO: descrypt token
	var requestToken *token
	if err := json.Unmarshal([]byte(requestTokenStr), requestToken); err != nil {
		return ""
	}
	return requestToken.ID
}
