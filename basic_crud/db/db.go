package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Value int    `json:"value" bson:"value"`
}

const (
	dialStr        = "localhost:27017"
	dbName         = "test"
	collectionName = "test"
)

var db *mgo.Database

func init() {
	session, err := mgo.Dial(dialStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB(dbName)
}

func collection() *mgo.Collection {
	return db.C(collectionName)
}

func GetAllItems() ([]Item, error) {
	res := []Item{}

	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func GetItem(id string) (*Item, error) {
	res := Item{}

	if err := collection().Find(bson.M{"_id": id}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func SaveItem(item Item) error {
	return collection().Insert(item)
}

func RemoveItem(id string) error {
	return collection().Remove(bson.M{"_id": id})
}
