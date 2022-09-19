package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"homework/MongoDB"
	"homework/server"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Pe4enka:GotIt1488@cluster0.ewt1q.mongodb.net/?retryWrites=true&w=majority"))
	MongoDB.Check(err)
	err = client.Connect(MongoDB.CTX)
	MongoDB.Check(err)
	err = client.Ping(MongoDB.CTX, nil)
	MongoDB.Check(err)
	fmt.Println("Connected to MongoDB!")
	MongoDB.EmployeeCollection = client.Database("employee").Collection("employee")

	service := server.Service{}
	service.Handler()

}
