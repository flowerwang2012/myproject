package main

import (
	"fmt"
	"flag"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
)
//mongodb复制集测试 mongo-go-driver驱动
//./mongo mongodb://smartgate:Smartgate2018@127.0.0.1:18809,127.0.0.1:18810,127.0.0.1:18811/admin
//use tif
//./mongoDriverTest --mongo "127.0.0.1:18809,127.0.0.1:18810,127.0.0.1:18811" --db_user "smartgate" --db_pwd "Smartgate2018" --op "insert/delete/select/update"
var (
	db = "tif"
	coll = "account"
)

type Account struct {
	UID     string `json:"uid" bson:"uid"`
	CID     string `json:"cid" bson:"cid"`
}

func main() {
	hosts := flag.String("mongo", "", "mongo url")
	dbUser := flag.String("db_user", "", "db user")
	dbPwd := flag.String("db_pwd", "", "db pwd")
	op := flag.String("op", "", "operator")
	flag.Parse()

	hostsStr := *hosts
	dbUserStr := *dbUser
	dbPwdStr := *dbPwd
	opStr := *op

	uri := fmt.Sprintf("mongodb://%s", hostsStr)
	if dbUserStr != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s/%s", dbUserStr, dbPwdStr, hostsStr, "admin")
	}
	fmt.Println(uri)
	client, err := mongo.NewClient(uri)
	if err != nil {
		fmt.Printf("mongo new client failed: %s \n", err.Error())
		return
	}
	err = client.Connect(nil)
	if err != nil {
		fmt.Printf("mongo client connect failed: %s", err.Error())
		return
	}
	mgo := client.Database(db)
	coll := mgo.Collection(coll)
	switch opStr {
	case "insert":
		account := &Account{
			UID: "12345678",
			CID: "445281199311077010",
		}
		doc, _ := bson.NewDocumentEncoder().EncodeDocument(account)
		result, err := coll.InsertOne(nil, doc)
		if err != nil {
			fmt.Printf("insert failed: %s \n", err.Error())
			return
		}
		fmt.Printf("insert seccess: %v \n", result.InsertedID)
	case "delete":
		result, err := coll.DeleteOne(nil, bson.NewDocument(bson.EC.String("cid", "445281199311077010")))
		if err != nil {
			fmt.Printf("delete failed: %s \n", err.Error())
			return
		}
		fmt.Printf("delete seccess: %d \n", result.DeletedCount)
	case "select":
		result := coll.FindOne(nil, bson.NewDocument(bson.EC.String("cid", "445281199311077010")))
		account := new(Account)
		if err := result.Decode(account); err != nil {
			fmt.Printf("select failed: %s \n", err.Error())
			return
		}
		fmt.Printf("select seccess: %+v \n", account)
	case "update":
		result, err := coll.UpdateOne(nil,
				bson.NewDocument(bson.EC.String("cid", "445281199311077010")),
				bson.NewDocument(
					bson.EC.SubDocumentFromElements("$set", bson.EC.String("name", "wanghongfa")),
				),
		)
		if err != nil {
			fmt.Printf("update failed: %s \n", err.Error())
			return
		}
		fmt.Printf("update seccess: %d \n", result.ModifiedCount)
	}

	//uri := fmt.Sprintf("mongodb://%s", "127.0.0.1:27017")
	//client, err := mongo.NewClient(uri)
	//if err != nil {
	//	fmt.Printf("mongo new client failed: %s \n", err.Error())
	//	return
	//}
	//err = client.Connect(nil)
	//if err != nil {
	//	fmt.Printf("mongo client connect failed: %s", err.Error())
	//	return
	//}
	//mgo := client.Database(db)
	//coll := mgo.Collection(coll)
	//result := coll.FindOne(nil, bson.NewDocument(bson.EC.String("cid", "445281199311077010")))
	//account := new(Account)
	//if err := result.Decode(account); err != nil {
	//	fmt.Printf("select failed: %s \n", err.Error())
	//	return
	//}
	//fmt.Printf("select seccess: %+v \n", account)
	fmt.Println("in the end")
}
