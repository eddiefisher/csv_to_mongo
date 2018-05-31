// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// product.go [created: Mon, 28 May 2018]

package product

import (
	"fmt"

	"github.com/globalsign/mgo"

	"github.com/globalsign/mgo/bson"
	"home.dev/toster/csv_to_mongo/src/dao"
)

const (
	COLLECTION = "products"
)

//Product Represents a product, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Product struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID   int           `bson:"product_id" json:"product_id"`
	Link        string        `bson:"link" json:"link"`
	Category    string        `bson:"category" json:"category"`
	Brand       string        `bson:"brand" json:"brand"`
	Name        string        `bson:"name" json:"name"`
	Fullname    string        `bson:"full_name" json:"full_name"`
	Specialname string        `bson:"special_name" json:"special_name"`
	Vendorname  string        `bson:"vendor_name" json:"vendor_name"`
	Description string        `bson:"description" json:"description"`
}

type Products []Product

func (xp *Products) BulkUpsertWithIndex() error {
	if err := createIndex(); err != nil {
		return err
	}

	bulk := dao.DB.C(COLLECTION).Bulk()
	bulk.Unordered()
	for _, product := range *xp {
		bulk.Upsert(bson.M{"product_id": product.ProductID}, product)
	}
	bulk.Upsert()
	_, err := bulk.Run()
	if err != nil {
		return err
	}

	return nil
}

func (p *Products) Count() (int, error) {
	return dao.DB.C(COLLECTION).Count()
}

func createIndex() error {
	collection := dao.DB.C(COLLECTION)
	err := collection.EnsureIndex(mgo.Index{Key: []string{"product_id"}})
	if err != nil {
		return err
	}
	indexes, err := collection.Indexes()
	if err != nil {
		return err
	}
	fmt.Println("key was created", indexes)

	return nil
}
