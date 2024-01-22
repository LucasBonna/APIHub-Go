package config

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

type Log struct {
	ID		 bson.ObjectId `bson:"_id,omitempty"`
	client_id: bson.ObjectId `bson:"_id,omitempty`
	request_id: bson.ObjectId `bson:"_id,omitempty`
	url: string `bson:"url"`
	method: string `bson:"string"`
	import "time"

	type Log struct {
		ID                   bson.ObjectId `bson:"_id,omitempty"`
		ClientID             bson.ObjectId `bson:"client_id,omitempty"`
		RequestID            bson.ObjectId `bson:"request_id,omitempty"`
		URL                  string        `bson:"url"`
		Method               string        `bson:"method"`
		StartRequest         time.Time     `bson:"start_request"`
		EndRequest           time.Time     `bson:"end_request"`
		ProcessTime          int           `bson:"process_time"`
		ResponseStatusCode   int           `bson:"response_status_code"`
	}
}

func CreateDefaultCollections(session *mgo.Session) error {
	db := session.DB("go")

	// Create User collection
	userCollection := db.C("users")
	err := userCollection.EnsureIndex(mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	})
	if err != nil {
		return err
	}

	// Create Logs collection
	postCollection := db.C("logs")
	err = postCollection.EnsureIndex(mgo.Index{
		Key:    []string{"title"},
		Unique: true,
	})
	if err != nil {
		return err
	}

	return nil
}
