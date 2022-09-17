package main

import(
	"fmt"
)

func main() {
	var conferenceName string = "Go conference"
	// conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainigTickets int = 50
	
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainigTickets is %T\n", conferenceName, conferenceTickets, remainigTickets )


	fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainigTickets, "are still available." )
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	//userName = "John"
	//userTickets = 2
	//fmt.Printf("User %v-%v booked %v tickets.\n", firstName, lastName, userTickets) 
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	
}