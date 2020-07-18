package dbs

import (
	"log"

	"github.com/globalsign/mgo"
)

var db *mgo.Database
var cname string = "product"

//Connecting <>
var Connecting *bool

func init() {
	Connecting = CheckifConnect()
	if !*Connecting {
		log.Fatal("not connect ")
		return
	}
	mgoInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "shop",
		Username: "michael",
		Password: "lneequal1",
	}
	session, err := mgo.DialWithInfo(mgoInfo)
	if err != nil {
		log.Fatalf("db error %s", err)
		panic(err)
	}
	db = session.DB("shop")
}

// CheckifConnect <>
func CheckifConnect() *bool {
	mgoInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "shop",
		Username: "michael",
		Password: "lneequal1",
	}
	session, err := mgo.DialWithInfo(mgoInfo)
	defer session.Close()
	result := err == nil
	return &result
}

// CreateProduct <this is a special case in this application>
func CreateProduct(inserData *Product) (*Product, error) {
	C := db.C(cname)
	err := C.Insert(&inserData)
	if err != nil {
		return nil, err
	}
	return inserData, nil
}

// ReadproductAll <>
func ReadproductAll() ([]Product, error) {
	entry := []Product{}
	C := db.C(cname)
	err := C.Find(nil).All(&entry)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

// // Create <>
// func create(collection string, data *interface{}) (*interface{}, error) {
// 	C := db.C(collection)
// 	err := C.Insert(&data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

// // ReadAll <>
// func readAll(collection string) ([]interface{}, error) {
// 	C := db.C(collection)
// 	var result []interface{}
// 	err := C.Find(nil).All(&result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// // UpdateWithID <>
// func updateWithID(collection, id string, updateData interface{}) (interface{}, error) {
// 	C := db.C(collection)
// 	err := C.UpdateId(id, &updateData)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return updateData, nil
// }
