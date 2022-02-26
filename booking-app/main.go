package main

import (
	"booking-app/helper"
	"fmt"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// user maps
// var bookings = make([]map[string]string, 0)

// use struct, mixed data type
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var bookings = make([]UserData, 0)

func main() {

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are sill available.")
	// fmt.Println("Get your tickets here to attend.")

	greetUsers()

	// var userName string
	// var userTickets int
	// userName = "Tom"
	// userTickets = 2
	// fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Array type: %T\n", bookings)
			// fmt.Printf("Array length: %v\n", len(bookings))

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("slice type: %T\n", bookings)
			// fmt.Printf("slice length: %v\n", len(bookings))

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
			// fmt.Println("Your input data is invalid, please try again!")
			if !isValidName {
				fmt.Println("first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid.")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are sill available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// use maps
		// firstNames = append(firstNames, booking["firstName"])
		// use struct
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// get input
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your fist name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter your tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// calculate
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	// bookings = append(bookings, firstName+" "+lastName)

	// create a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// use struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
