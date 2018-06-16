package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/big"
	"os"

	"../constants"
	"../log"
)

type token struct {
	ID       string `json:"id,omitempty"`
	Salt     int64  `json:"salt,omitempty"`
	LifeTime int    `json:"life,omitempty"`
}

const maxInt = int64((^uint64(0)) >> 1)

func GenerateToken(id string) string {
	salt, _ := rand.Int(rand.Reader, big.NewInt(maxInt))
	newToken := &token{ID: id, Salt: salt.Int64()}

	jsonToken, _ := json.Marshal(newToken)

	return encryptToken(jsonToken)
}

func GetIdFromToken(requestTokenStr string) string {

	requestToken := &token{}
	requestTokenStr = decryptToken(requestTokenStr)
	if err := json.Unmarshal([]byte(requestTokenStr), requestToken); err != nil {
		return ""
	}
	return requestToken.ID
}

func encryptToken(clearToken []byte) string {
	tokenBytes := make([]byte, aes.BlockSize+len(clearToken))
	iv := tokenBytes[:aes.BlockSize]

	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(tokenBytes[aes.BlockSize:], clearToken)

	return base64.URLEncoding.EncodeToString(tokenBytes)
}

func decryptToken(encryptedToken string) string {
	decodedToken, _ := base64.URLEncoding.DecodeString(encryptedToken)

	if len(decodedToken) < aes.BlockSize {
		return ""
	}
	iv := decodedToken[:aes.BlockSize]
	tokenBytes := decodedToken[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(tokenBytes, tokenBytes)

	return string(tokenBytes)
}

var block cipher.Block

func CreateSecret() {

	if block == nil {
		var key = make([]byte, constants.SecretSize)
		keyBytes, err := ioutil.ReadFile(constants.SecretFileName)
		if err != nil {

			rndBytes := make([]byte, constants.SecretSize)
			keyBytes = make([]byte, hex.EncodedLen(constants.SecretSize))

			_, err := rand.Read(rndBytes)

			log.Log(constants.MessageSecretCreatingFile)

			f, err := os.Create(constants.SecretFileName)
			if err != nil {
				log.LogNow(constants.ErrorFileSecret + err.Error())
				panic(err)
			}

			hex.Encode(keyBytes, rndBytes)
			_, err = f.Write(keyBytes)
			if err != nil {
				log.LogNow(constants.ErrorFileSecret + err.Error())
				panic(err)
			}

			err = f.Sync()
			if err != nil {
				log.LogNow(constants.ErrorFileSecret + err.Error())
				panic(err)
			}
		}

		_, err = hex.Decode(key, keyBytes)
		if err != nil {
			log.LogNow(constants.ErrorDecodingSecret + err.Error())
			panic(err)
		}

		block, err = aes.NewCipher(key)
		if err != nil {
			panic(err.Error())
		}
	}

}
