package dbs

import "github.com/globalsign/mgo/bson"

// Product <>
type Product struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Title   string        `bson:"title"`
	Price   float64       `bson:"price"`
	Code    string        `bson:"code"`
	Content string        `bson:"content"`
}
