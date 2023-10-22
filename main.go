package main

import (
	"fmt"
	"strings"
)

func main() {

	var app = "Barbie Movie"
	var tickets uint = 50
	var bookings []string

	fmt.Println("Welcome to the", app)
	fmt.Printf("We have %v tickets remaining.\n", tickets)

	for {
		var firstName string

		var lastName string

		var userTickets uint

		fmt.Print("Enter your first name: ")

		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")

		fmt.Scan(&lastName)

		fmt.Printf("Hello %v, how many tickets would you like to buy?\n", firstName)

		fmt.Scan(&userTickets)

		if tickets == 0 {
			fmt.Printf("Sorry %v, we are sold out.\n", firstName)
			continue
		} else if userTickets > tickets {
			fmt.Printf("Sorry %v, we only have %v tickets left.\n", firstName, tickets)
			continue
		} else {
			tickets -= userTickets
			firstNames := []string{}
			bookings = append(bookings, firstName)

			fmt.Printf("Thank you %v, you have purchased %v tickets.\n", firstName, userTickets)
			fmt.Printf("We have %v tickets remaining.\n", tickets)

			for _, booking := range bookings {
				names := strings.Split(booking, " ")
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The following people have purchased tickets: %v\n", firstNames)
		}

	}

}
