package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets = conferenceTickets
	var bookings []string

	// fmt.Printf("conferenceName is %T\n", conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var bookings = [50]string{}
	// bookings[0] = "Michall"

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask for their name
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The wole array: %v\n", bookings)
		fmt.Printf("The wole array length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings {
			var name = strings.Fields(booking)[0]
			firstNames = append(firstNames, name)
		}

		fmt.Printf("The First names of bookings are: %v\n", firstNames)
	}
}
