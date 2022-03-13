package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var client *mongo.Database

func initMongo()  {
	var err error
	uri:= "mongodb://admin:123456@localhost:27017"
	client, err = ConnectToDB(uri, "testing", 2*time.Second, 10)
	if err != nil {
		log.Fatalln(err)
	}
}

func mongoAccess() {

	collection := client.Collection("numbers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		log.Println(err)
		return
	}
	id := res.InsertedID
	log.Println("insertedID:", id)
}

// pool 连接池模式
func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error)  {
	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// 通过传进来的uri连接相关的配置
	o := options.Client().ApplyURI(uri)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	o.SetMaxPoolSize(num)
	// 发起链接
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func() {
		//client.Disconnect(ctx)
	}()

	// 判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 返回 client
	return client.Database(name), nil
}