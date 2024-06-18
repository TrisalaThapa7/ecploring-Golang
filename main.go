package main

import (
	"fmt"
	"strconv"
)

var conference_name string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) //creating an empty slice of maps/creating empty list of maps

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames() //function
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

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application.\n", conference_name)
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T. \n", conference_name, conferenceTickets, remainingTickets)

	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking["firstNames"])
	}

	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = make(map[string]string) // can be of different type. however we have key string and value also string. // We cannot mix data type!!! // when we created a map we have a type of the map and we wrap it to make function and that creates an empty map.
	// here creating an empty map.
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // for userTickets: as it is uint data type, we have to convert uint into a string.

	// FormatUnit function takes our uint value which maybe anything between 1 to 50 and formats it to a string as a decimal number. And 10 is for base 10 which represents decimal numbers.

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conference_name)
}
