package common

import (
	"gopkg.in/mgo.v2"
)

const (
	MONGO_URL = "mongodb://127.0.0.1:27017"
	// WUSHU_DB = "wushu"
)

func mgoSession() *mgo.Session {
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

func DB(dbname string) *mgo.Database {
	db := mgoSession().DB(dbname)
	return db
}

func MgoClose() {
	mgoSession().Close()
}
