// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// main.go [created: Mon, 28 May 2018]

package main

import (
    "fmt"
	"home.dev/toster/csv_to_mongo/src/version"
)

func main() {
    fmt.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
}
