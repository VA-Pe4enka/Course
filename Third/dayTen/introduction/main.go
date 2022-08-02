package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	Name      string   `bson:"name"`
	Year      string   `bson:"year"`
	Directors []string `bson:"directors"`
	Writers   []string `bson:"writers"`
	BoxOffice `bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross  uint64 `bson:"gross"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB successfully")

	collection := client.Database("introDB").Collection("movies")

	firstMovie := Movie{
		Name:      "First Movie",
		Year:      "2020",
		Directors: []string{"Some User"},
		Writers:   []string{"User One", "User Two"},
		BoxOffice: BoxOffice{
			Budget: 188000,
			Gross:  533000,
		},
	}

	_, err = collection.InsertOne(context.TODO(), firstMovie)

	if err != nil {
		log.Fatal(err)
	}

	queryResult := &Movie{}
	// bson.M should use for nested fields
	filter := bson.M{"boxOffice.budget": bson.M{"$gt": 150000}}
	result := collection.FindOne(context.TODO(), filter)
	err = result.Decode(queryResult)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie:", queryResult)
	err = client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Disconnected from MongoDB")
}
