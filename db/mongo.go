package db

import (
	mgo "gopkg.in/mgo.v2"
)

// DB is mongodb database
var DB *mgo.Database

// Connect is connection to database
func Connect() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	DB = session.DB("gographql")
}

func init() {
	Connect()
}
