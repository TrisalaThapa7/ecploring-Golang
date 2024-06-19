package main

import (
	"fmt"
	"sync"
	"time"
)

var conference_name string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0) // creates or initiates empty list of user data structs

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)

		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames() //function
		fmt.Printf("The first names of bookings are: %v \n", firstNames)

		if remainingTickets == 0 {

			fmt.Println("Our conference is fully booked, please come next time!")
			// break
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

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application.\n", conference_name)
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T. \n", conference_name, conferenceTickets, remainingTickets)

	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
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

	var userData = UserData{

		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conference_name)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########")
	fmt.Printf("Sending ticket: %v \n to email address %v", ticket, email)
	fmt.Println("########")

	wg.Done()
}
