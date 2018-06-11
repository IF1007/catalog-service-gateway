package db

import (
	"sort"

	"../constants/"
	"../log/"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserAuth struct {
	Login string `json:"login,omitempty"`
	Pass  string `json:"password,omitempty"`
	ID    string
}

const databaseName = "csdAuth"
const collectionUser = "userAuth"

var userCollection *mgo.Collection

func Start() {
	log.LogNow(constants.MessageTryConnectDB)
	session, err := mgo.Dial(constants.MongoDBHost)
	if err != nil {
		log.LogNow(constants.ErrorTryingConnectDB)
		log.LogNow(err.Error())
		panic(err)
	}
	log.Log(constants.MessageConnectDBSuccess)
	db := session.DB(databaseName)
	defer session.Close()
	collections, err := db.CollectionNames()
	needCreateCollection := sort.SearchStrings(collections, collectionUser) > 0
	userCollection = db.C(collectionUser)

	if needCreateCollection {
		userCollection.Create(&mgo.CollectionInfo{DisableIdIndex: true})
	}
}

func InsertUser(user *UserAuth) (bool, error) {
	if user.Login == "" || user.Pass == "" {
		return true, nil
	}
	user.ID = uuid.Must(uuid.NewV4()).String()
	newUsr := bson.M{"Login": user.Login, "Password": user.Pass, "ID": user.ID}
	err := userCollection.Insert(newUsr)

	return false, err
}

func GetUserByID(id string) (UserAuth, error) {
	var result UserAuth
	err := userCollection.FindId(id).One(&result)
	return result, err
}

func GetUserByLoginPass(login string, pass string) (*UserAuth, error) {
	var result *UserAuth

	if login == "" || pass == "" {
		return nil, nil
	}

	hasLogin, err := HasLogin(login)
	if !hasLogin || err != nil {
		return nil, nil
	}

	params := bson.M{"Login": login, "Password": pass}
	err = userCollection.Find(params).One(&result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func HasLogin(login string) (bool, error) {
	if login == "" {
		return false, nil
	}
	n, err := userCollection.Find(bson.M{"Login": login}).Count()
	return n > 0, err
}
