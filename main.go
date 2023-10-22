package main

import (
	"fmt"
)

var app = "Barbie Movie"
var tickets uint = 50
var bookings []string

func main() {

	greet()

	for tickets > 0 && len(bookings) < 50 {
		firstName, _, boughtTickets := getUserInput()

		if boughtTickets > tickets {
			fmt.Printf("Sorry %v, we only have %v tickets left.\n", firstName, tickets)
			continue
		} else {
			tickets -= boughtTickets
			bookings = append(bookings, firstName)

			fmt.Printf("Thank you %v, you have purchased %v tickets.\n", firstName, boughtTickets)
			fmt.Printf("We have %v tickets remaining.\n", tickets)

			fmt.Printf("The following people have purchased tickets: %v\n", getFirstNames())
		}

	}

	fmt.Printf("Sorry, we are sold out.\n")

}
