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

func mongoAccess() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("connect ok")

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("ping ok")

	collection := client.Database("testing").Collection("numbers")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		log.Println(err)
		return
	}
	id := res.InsertedID
	log.Println("insertedID:", id)

}
