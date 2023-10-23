package main

import (
	"fmt"
	"strconv"
)

var app = "Barbie Movie"
var tickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greet()

	for tickets > 0 && len(bookings) < 50 {
		firstName, lastName, boughtTickets := getUserInput()

		if boughtTickets > tickets {
			fmt.Printf("Sorry %v, we only have %v tickets left.\n", firstName, tickets)
			continue
		} else {
			tickets -= boughtTickets

			var userData = make(map[string]string)

			userData["firstName"] = firstName
			userData["lastName"] = lastName
			userData["boughtTickets"] = strconv.FormatUint(uint64(boughtTickets), 10)

			bookings = append(bookings, userData)

			fmt.Printf("Thank you %v, you have purchased %v tickets.\n", firstName, boughtTickets)
			fmt.Printf("We have %v tickets remaining.\n", tickets)
			fmt.Printf("The following people have purchased tickets: %v\n", getFirstNames())

		}

	}

	fmt.Printf("Sorry, we are sold out.\n")

}
