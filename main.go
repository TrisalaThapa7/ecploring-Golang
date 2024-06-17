package main

import (
	"fmt"
	"strings"
)

func main() {
	var conference_name string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conference_name, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conference_name)

			firstNames := getFirstNames(bookings) //function
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {

				fmt.Println("Our conference is fully booked, please come next time!")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("firstname or lastname you entered is short.")
			}
			if !isValidEmail {
				fmt.Println("email address you've entered is not valid.")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets entered is invalid.")
			}

		}

	}

}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application.\n", confName)
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T. \n", confName, confTickets, remainingTickets)

	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, confName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v conferenceName.\n This is list %v", remainingTickets, confName, bookings)
}
