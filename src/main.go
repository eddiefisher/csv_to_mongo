// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// main.go [created: Mon, 28 May 2018]

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"home.dev/toster/csv_to_mongo/src/dao"
	"home.dev/toster/csv_to_mongo/src/mgo/product"
)

func init() {
	dao.Connect()
}

func main() {
	csvFile, err := os.Open("./products_ru.csv")
	if err != nil {
		_ = fmt.Errorf("could not open file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var prod product.Product
	var xprod product.Products

	for _, data := range csvData {
		prod.ProductID, _ = strconv.Atoi(data[0])
		prod.Link = data[1]
		prod.Category = data[2]
		prod.Brand = data[3]
		prod.Name = data[4]
		prod.Fullname = data[5]
		prod.Specialname = data[6]
		prod.Vendorname = data[7]
		prod.Description = data[8]
		xprod = append(xprod, prod)
	}

	err = xprod.Upsert()
	if err != nil {
		log.Fatalf("could not insert or update: %v", err)
	}
	i, _ := xprod.Count()
	fmt.Println("Count of record:", i)
}
