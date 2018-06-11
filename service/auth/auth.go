package auth

import "math/rand"

type token struct {
	ID   string `json:"id,omitempty"`
	Salt int    `json:"salt,omitempty"`
}

const maxInt = int((^uint(0)) >> 1)

func GenerateToken(id string) string {
	newToken := &token{ID: id, Salt: rand.Intn(maxInt)}
	return "token:" + string(newToken.Salt)
}

func IsTokenValid(token string) bool {
	return true
}
