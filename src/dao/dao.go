package dao

import (
	"fmt"

	"github.com/globalsign/mgo"
	"home.dev/toster/csv_to_mongo/src/config"
)

var DB *mgo.Database

func Connect() {
	Host := []string{
		config.MongoHost,
	}

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Database: config.MongoDatabase,
		// Username:       config.MongoUsername,
		// Password:       config.MongoPassword,
		// ReplicaSetName: config.MongoReplicaSetName,
	})
	if err != nil {
		panic(err)
	}

	DB = session.DB(config.MongoDatabase)
	fmt.Printf("Connected to replica set %v!\n", session.LiveServers())
}
