// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// dao.go [created: Mon, 28 May 2018]

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
