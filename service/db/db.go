package db

import (
	"sort"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"../constants"
	"../log"
	"github.com/satori/go.uuid"
)

type UserAuth struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	ID       string
}

const databaseName = "csdAuth"
const collectionUser = "userAuth"

var userCollection *mgo.Collection

func Start() {
	log.LogNow(constants.MessageTryConnectDB)
	// TODO: add auth in DB
	session, err := mgo.Dial(constants.HostMongoDB)
	if err != nil {
		log.LogNow(constants.ErrorTryingConnectDB)
		log.LogNow(err.Error())
		panic(err)
	}
	log.Log(constants.MessageConnectDBSuccess)
	db := session.DB(databaseName)
	// defer session.Close()
	collections, err := db.CollectionNames()
	needCreateCollection := sort.SearchStrings(collections, collectionUser) > 0
	userCollection = db.C(collectionUser)

	if needCreateCollection {
		userCollection.Create(&mgo.CollectionInfo{DisableIdIndex: true})
	}
}

func InsertUser(user *UserAuth) (bool, error) {
	if user.Login == "" || user.Password == "" {
		return true, nil
	}
	user.ID = uuid.Must(uuid.NewV4()).String()
	newUsr := bson.M{"Login": user.Login, "Password": user.Password, "ID": user.ID}
	err := userCollection.Insert(&newUsr)

	return false, err
}

func GetUserByID(id string) (UserAuth, error) {
	var result UserAuth
	err := userCollection.FindId(id).One(&result)
	return result, err
}

func GetUserByLoginPass(login string, pass string) (*UserAuth, error) {

	if login == "" || pass == "" {
		return nil, nil
	}

	hasLogin, err := HasLogin(login)
	if !hasLogin || err != nil {
		return nil, nil
	}

	params := bson.M{"Login": login, "Password": pass}

	jsonObj := bson.M{}
	err = userCollection.Find(params).One(&jsonObj)

	if err != nil {
		return nil, err
	}

	return &UserAuth{
		ID: jsonObj["ID"].(string),
	}, err
}

func HasLogin(login string) (bool, error) {
	if login == "" {
		return false, nil
	}
	n, err := userCollection.Find(bson.M{"Login": login}).Count()
	return n > 0, err
}
