package main

import (
	"fmt"
	"time"
)

func greet() {
	fmt.Printf("Hello, welcome to the %v\n", app)
	fmt.Printf("We have %v tickets remaining.\n", tickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
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

func sendTicket(boughtTickets uint, firstName string, lastName string) {
	var message = fmt.Sprintf("%v tickets for %v %v", boughtTickets, firstName, lastName)

	time.Sleep(5 * time.Second)
	fmt.Printf("Sending ticket...\n")

	time.Sleep(1 * time.Second)
	fmt.Printf("Successfully sent %v\n", message)

	wg.Done()
}
