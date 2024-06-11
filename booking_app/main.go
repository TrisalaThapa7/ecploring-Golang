package main

import "fmt"

func main() {
	var conference_name = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conference_name is %T, conferenceTickets is %T, remainingTickets is %T. \n", conference_name, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to our %v booking application.\n", conference_name)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend.")

	var userName string
	var userTickets int

	userName = "Seema"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
