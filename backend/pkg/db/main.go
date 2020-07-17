package dbs

import (
	"log"

	"github.com/globalsign/mgo"
)

var db *mgo.Database

func init() {
	mgoInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "shop",
		Username: "michael",
		Password: "lneequal1",
	}
	session, err := mgo.DialWithInfo(mgoInfo)
	if err != nil {
		log.Fatalf("db error %s", err)
	}
	db = session.DB("shop")
}

// Create <>
func Create(collection string, data *interface{}) (*interface{}, error) {
	C := db.C(collection)
	err := C.Insert(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ReadAll <>
func ReadAll(collection string) ([]interface{}, error) {
	C := db.C(collection)
	var result []interface{}
	err := C.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateWithID <>
func UpdateWithID(collection, id string, updateData interface{}) (interface{}, error) {
	C := db.C(collection)
	err := C.UpdateId(id, &updateData)
	if err != nil {
		return nil, err
	}
	return updateData, nil
}
