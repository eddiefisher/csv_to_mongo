package dao

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"home.dev/toster/csv_to_mongo/src/config"
)

var DB *mgo.Database

//Connect info https://godoc.org/github.com/globalsign/mgo#DialInfo
func Connect() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{config.MongoHost},
		Database: config.MongoDatabase,
		Timeout:  10 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	DB = session.DB(config.MongoDatabase)
	fmt.Printf("Connected to replica set %v!\n", session.LiveServers())
}
