[![Go Report Card](https://goreportcard.com/badge/github.com/eddiefisher/csv_to_mongo)](https://goreportcard.com/report/github.com/eddiefisher/csv_to_mongo)

[godoc.org]: http://godoc.org/home.dev/toster/csv_to_mongo "godoc.org"

## Install

    go get github.com/eddiefisher/csv_to_mongo

## Docs

Tested on CSV file with 70500 records

- create or update: 20 min
- upsert without index: 29 min
- upsert with index: 51 sec
- bulk upsert with index: ~9 sec

## Author

Eddie Fisher [eddi.fisher@gmail.com]

## Copyright & License

Copyright (c) 2018, Eddie Fisher.
All rights reserved.
Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.
