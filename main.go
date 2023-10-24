package main

import (
	"fmt"
	"sync"
)

var app = "Barbie Movie"
var tickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName     string
	lastName      string
	boughtTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greet()

	firstName, lastName, boughtTickets := getUserInput()

	if boughtTickets > tickets {
		fmt.Printf("Sorry %v, we only have %v tickets left.\n", firstName, tickets)
	} else {
		tickets -= boughtTickets

		var userData = UserData{
			firstName:     firstName,
			lastName:      lastName,
			boughtTickets: boughtTickets,
		}

		bookings = append(bookings, userData)

		fmt.Printf("Thank you %v, you have purchased %v tickets.\n", firstName, boughtTickets)
		fmt.Printf("We have %v tickets remaining.\n", tickets)
		fmt.Printf("The following people have purchased tickets: %v\n", getFirstNames())

		wg.Add(1)
		go sendTicket(boughtTickets, firstName, lastName)

		if tickets == 0 {
			fmt.Printf("Sorry, we are sold out!\n")
		}

	}
	wg.Wait()
}
