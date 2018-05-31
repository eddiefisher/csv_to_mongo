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
	defer dao.DB.Session.Close()

	data, err := readCSV()
	if err != nil {
		log.Fatalln(err)
	}

	var prod product.Product
	var xprod product.Products

	for _, data := range data {
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

	err = xprod.BulkUpsertWithIndex()
	if err != nil {
		log.Fatalf("could not insert or update: %v", err)
	}
	i, _ := xprod.Count()
	fmt.Println("Count of record:", i)
}

func readCSV() ([][]string, error) {
	file, err := os.Open("./products_ru.csv")
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}
	return data, nil
}
