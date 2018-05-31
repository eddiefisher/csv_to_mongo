// https://gist.github.com/boj/5412538
// https://codereview.stackexchange.com/questions/70274/parsing-csvs-for-bulk-database-insertions

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

func (xp *Products) FindAll() error {
	return dao.DB.C(COLLECTION).Find(nil).All(xp)
}

func (xp *Products) Upsert() error {
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

	for _, product := range *xp {
		_, err := collection.Upsert(bson.M{"product_id": product.ProductID}, product)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Product) FindById(id string) error {
	return dao.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&p)
}

func (p *Product) Insert(product Product) error {
	return dao.DB.C(COLLECTION).Insert(&product)
}

func (p *Product) Delete(product Product) error {
	return dao.DB.C(COLLECTION).Remove(&product)
}

func (p *Product) Update(product Product) error {
	return dao.DB.C(COLLECTION).UpdateId(product.ID, &product)
}

func (p *Products) Count() (int, error) {
	return dao.DB.C(COLLECTION).Count()
}
