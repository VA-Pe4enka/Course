package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var ScheduleCollection *mongo.Collection
var StationCollection *mongo.Collection
var ctx = context.TODO()

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
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

	ScheduleCollection = client.Database("Cluster0").Collection("schedules")
	StationCollection = client.Database("Cluster0").Collection("stations")
}

func postStation() {
	var station Station

	fmt.Print("Enter station id: ")
	fmt.Fscan(os.Stdin, &station.ID)
	fmt.Println()
	fmt.Print("Enter station opening time: ")
	fmt.Fscan(os.Stdin, &station.OpeningTime)
	fmt.Println()
	fmt.Print("Enter station closing time: ")
	fmt.Fscan(os.Stdin, &station.ClosinTime)

	insertResult, err := StationCollection.InsertOne(ctx, station)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func postSchedule() {
	var schedule Schedule

	fmt.Print("Enter station id: ")
	fmt.Fscan(os.Stdin, &schedule.StationID)
	fmt.Println()
	fmt.Print("Enter schedule id: ")
	fmt.Fscan(os.Stdin, &schedule.Num)
	fmt.Println()
	fmt.Print("Enter direction: ")
	fmt.Fscan(os.Stdin, &schedule.Direction)
	fmt.Println()
	fmt.Print("Enter arrival time: ")
	fmt.Fscan(os.Stdin, &schedule.ArrivalTime)
	fmt.Println()
	fmt.Print("Enter departure time: ")
	fmt.Fscan(os.Stdin, &schedule.DepartureTime)

	insertResult, err := ScheduleCollection.InsertOne(ctx, schedule)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func getStation(id string) (Station, []*Schedule) {

	var station Station
	var schedule []*Schedule

	filter := bson.D{{"id", id}}
	err := StationCollection.FindOne(ctx, filter).Decode(&station)
	if err != nil {
		log.Fatal(err)
	}

	opt := options.Find()
	newFilter := bson.M{"stationid": id}
	cur, err := ScheduleCollection.Find(ctx, newFilter, opt)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var element Schedule
		err := cur.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}

		schedule = append(schedule, &element)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", schedule)

	return station, schedule

}

func deleteSchedule(id string) {

	filter := bson.M{"stationid": id}
	deleteResult, err := ScheduleCollection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the schedules collection\n", deleteResult.DeletedCount)
}
