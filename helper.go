package main

import (
	"fmt"
)

func greet() {
	fmt.Printf("Hello, welcome to the %v\n", app)
	fmt.Printf("We have %v tickets remaining.\n", tickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func getUserInput() (string, string, uint) {
	var firstName string
	var lastName string
	var boughtTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Hello %v, how many tickets would you like to buy?\n", firstName)
	fmt.Scan(&boughtTickets)

	return firstName, lastName, boughtTickets
}
