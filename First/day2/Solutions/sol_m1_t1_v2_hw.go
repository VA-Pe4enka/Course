package main

import (
	"fmt"
	"log"
)

const costPetrolRUB int = 48

func main() {
	var distance int

	var priceOfTrip int

	fmt.Printf("Enter distance: ")
	fmt.Scanf("%d\n", &distance)
	fmt.Printf("%d\n", distance)

	var petrolMileage int
	fmt.Println("Enter rashod v litrah: ")
	fmt.Scanf("%d\n", &petrolMileage)
	fmt.Printf("%d\n", petrolMileage)

	if distance < 50 || distance > 10000 {
		fmt.Println("Wrong distance!")
		log.Panic()
	} else if petrolMileage < 4 || petrolMileage > 25 {
		fmt.Println("Wrong petrol rashod!")
		log.Panic()
	} else {
		priceOfTrip = (distance / 100) * petrolMileage * costPetrolRUB
		fmt.Println("Price of trip = ", priceOfTrip)
	}

}
