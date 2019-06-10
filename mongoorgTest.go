package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"time"
)

type UserInfo struct {
	Name      string `json:"name" bson:"name"`
	CID       string `json:"cid" bson:"cid"`
	Starttime string `json:"starttime" bson:"starttime"`
	Endtime   string `json:"endtime" bson:"endtime"`
}

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	/*** code begin ***/

	uri := fmt.Sprintf("mongodb://%s", "127.0.0.1:27017")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	coll := client.Database("test").Collection("user")
	user := new(UserInfo)
	filter := bson.M{"cid": "445281199311077010"}

	result := coll.FindOne(ctx, filter)
	result.Decode(user)

	user.Name = "wanghongfa"
	_, err = coll.UpdateOne(ctx, filter, bson.M{"$set": user})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", user)

	/*** code end ***/
	select {
	case <-ch:
		fmt.Println("退出程序")
		os.Exit(2)
	}
}
