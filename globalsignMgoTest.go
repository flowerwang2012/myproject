package main

import (
	"fmt"
	"flag"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)
//mongodb复制集测试 globalsign/mgo驱动
//./mongo mongodb://smartgate:Smartgate2018@127.0.0.1:18809,127.0.0.1:18810,127.0.0.1:18811/admin
//use tif
//./globalsignMgoTest --mongo "127.0.0.1:18809,127.0.0.1:18810,127.0.0.1:18811" --db_user "smartgate" --db_pwd "Smartgate2018" --op "insert/delete/select/update"
var (
	db1 = "tif"
	coll1 = "account"
	mgoSession *mgo.Session
)

type Account1 struct {
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

	url := fmt.Sprintf("mongodb://%s", hostsStr)
	if dbUserStr != "" {
		url = fmt.Sprintf("mongodb://%s:%s@%s/%s", dbUserStr, dbPwdStr, hostsStr, "admin")
	}
	fmt.Println(url)
	var err error
	mgoSession, err = mgo.Dial(url)
	if err != nil {
		return
	}
	mgoSession.SetMode(mgo.Monotonic, true)
	switch opStr {
	case "insert":
		account := &Account1{
			UID: "12345678",
			CID: "445281199311077010",
		}
		err = MC(coll1, func(coll *mgo.Collection) error {
			err = coll.Insert(account)
			return err
		})
		if err != nil {
			fmt.Printf("insert failed: %s \n", err.Error())
			return
		}
		fmt.Printf("insert seccess \n")
	case "delete":
		doc := bson.M{"cid": "445281199311077010"}
		err = MC(coll1, func(coll *mgo.Collection) error {
			err = coll.Remove(doc)
			return err
		})
		if err != nil {
			fmt.Printf("delete failed: %s \n", err.Error())
			return
		}
		fmt.Printf("delete seccess \n")
	case "select":
		account := new(Account1)
		doc := bson.M{"cid": "445281199311077010"}
		err = MC(coll1, func(coll *mgo.Collection) error {
			err = coll.Find(doc).One(account)
			return err
		})
		if err != nil {
			fmt.Printf("select failed: %s \n", err.Error())
			return
		}
		fmt.Printf("select seccess: %+v \n", account)
	case "update":
		doc := bson.M{"cid": "445281199311077010"}
		update := bson.M{"$set": bson.M{"name": "wanghongfa"}}
		err = MC(coll1, func(coll *mgo.Collection) error {
			err = coll.Update(doc, update)
			return err
		})
		if err != nil {
			fmt.Printf("update failed: %s \n", err.Error())
			return
		}
		fmt.Printf("update seccess \n", )
	}
	fmt.Println("in the end")
}

// 可以指定collection
func MC(collection string, f func(*mgo.Collection) error) error {
	sessionCopy := mgoSession.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(db1).C(collection)
	return f(c)
}

// 可以指定database和collection
func MDC(dbName string, collection string, f func(*mgo.Collection) error) error {
	sessionCopy := mgoSession.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(dbName).C(collection)
	return f(c)
}