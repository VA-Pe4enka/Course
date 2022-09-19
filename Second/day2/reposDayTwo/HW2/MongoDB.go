package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var TimeCollection *mongo.Collection
var UsersCollection *mongo.Collection
var ctx = context.TODO()

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.ewt1q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	TimeCollection = client.Database("Cluster0").Collection("time")
	UsersCollection = client.Database("Cluster0").Collection("users")
}

func getEmployee(id string) Person {

	var person Person

	filter := bson.M{"id": id}
	err := UsersCollection.FindOne(ctx, filter).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", person)

	return person
}

func postEmployee() {

	var person Person

	fmt.Print("Enter user id: ")
	fmt.Fscan(os.Stdin, &person.ID)
	fmt.Println()
	fmt.Print("Enter user full name: ")
	fmt.Fscan(os.Stdin, &person.FIO)
	fmt.Println()
	fmt.Print("Enter user department: ")
	fmt.Fscan(os.Stdin, &person.DEPARTMENT)
	fmt.Println()
	fmt.Print("Enter user position: ")
	fmt.Fscan(os.Stdin, &person.POSITION)
	fmt.Println()
	person.WORKTIME = 0

	insertResult, err := UsersCollection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func createNote(id string) Note {
	var note Note

	note.ID = id
	note.StartTime = time.Now()

	insertResult, err := TimeCollection.InsertOne(ctx, note)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return note
}

func updateNote(id string) {
	var note Note
	filter := bson.M{"id": id}
	err := TimeCollection.FindOne(ctx, filter).Decode(&note)
	if err != nil {
	}

	note.EndTime = time.Now()
	note.Result = int(note.EndTime.Sub(note.StartTime) / time.Minute)

	filter = bson.M{"id": id}
	deleteResult, err := TimeCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	insertResult, err := TimeCollection.InsertOne(ctx, note)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}

func updateEmployee(id string) {
	var note Note
	var person Person

	filter := bson.M{"id": id}
	err := UsersCollection.FindOne(context.TODO(), filter).Decode(&person)
	if err != nil {
	}
	filter = bson.M{"id": id}
	err = TimeCollection.FindOne(context.TODO(), filter).Decode(&note)
	if err != nil {
	}

	person.WORKTIME = +note.Result

	deleteResult, err := UsersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	insertResult, err := UsersCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}
