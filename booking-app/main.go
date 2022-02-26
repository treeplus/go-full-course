package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are sill available.")
	// fmt.Println("Get your tickets here to attend.")

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are sill available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// var userName string
	// var userTickets int
	// userName = "Tom"
	// userTickets = 2
	// fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

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

	// calculate
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
