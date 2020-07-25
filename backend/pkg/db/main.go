package dbs

// TODO:https://studygolang.com/articles/14152 研究一下這個問題解法
import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// InsertInChan <>
var InsertInChan chan Product

// InsertOutChan <>
var InsertOutChan chan Product

// ShopDB <>
var shop *mgo.Database

// ConnectDB <>
func ConnectDB() {
	mgoInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "shop",
		Username: "michael",
		Password: "lneequal1",
	}
	session, err := mgo.DialWithInfo(mgoInfo)
	if err != nil {
		log.Fatalf("db error : %s", err)
		panic("db connect error")
	}
	shop = session.DB("shop")
}

func getCollection(name string) *mgo.Collection {
	if shop != nil {
		return shop.C(name)
	}
	return nil
}

// SetinsertInChan <>
func SetinsertInChan(data Product) {
	InsertInChan <- data
}

// GetinsertOutChan <>
func GetinsertOutChan() <-chan Product {
	return InsertOutChan
}

// Insertprocess <>
func Insertprocess() {
	go func(in *chan Product) {
		for {
			select {
			case input := <-*in:
				C := getCollection("product")
				if C == nil {
					panic("db process error")
				}
				err := C.Insert(&input)
				if err != nil {
					log.Fatalf("insert error %s", err)
					panic("insert error")
				}
				oData := Product{}
				err = C.Find(bson.M{"title": input.Title}).One(&oData)
				if err != nil {
					log.Fatalf("insert return error %v", err)
					panic("insert return error")
				}
				InsertOutChan <- oData
			}
		}
	}(&InsertInChan)
}
