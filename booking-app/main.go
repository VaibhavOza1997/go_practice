package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) // slice of maps, list of maps
//var bookings = []string{} // - slice asignment with no size and empty values also called dynamic array

// firstName : "Vaibhav"
// lastName := "Oza"
// email := "vo@g.com"
// userTickets := 2

func main() {
	// conferenceName := "Go Conference"
	// const conferenceTickets = 50
	// var remainingTickets uint = 50

	//var bookings = [50]string{} - array asignment with size and empty values
	//var bookings = [50]string{"Nana", "Nicole", "Peter"} - array asignment with size and initial values
	//var bookings [50]string // - array asignment with size and empty values
	//var bookings []string // - slice asignment with no size and empty values also called dynamic array
	//var bookings = []string{} // - slice asignment with no size and empty values also called dynamic array
	bookings := make([]map[string]string, 0) // slice of maps, list of maps

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	// fmt.Printf(("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n"), conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Hello, World!")

	//conferenceTickets=30
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	// fmt.Println("Get your tickets to attend!!!")

	for {

		firstName, lastName, email, userTickets := getUserInput()
		// var firstName string
		// var lastName string
		// var email string
		// var userTickets uint

		//ask user for their name
		// fmt.Println("Enter your first name: ")
		// fmt.Scan(&firstName)

		// fmt.Println("Enter your last name: ")
		// fmt.Scan(&lastName)

		// fmt.Println("Enter your email: ")
		// fmt.Scan(&email)

		// //ask user for number of tickets
		// fmt.Println("Enter number of tickets: ")
		// fmt.Scan(&userTickets)

		// if userTickets > remainingTickets {
		// 	fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		// 	//break //exit the program
		// 	continue //continue to the next iteration of the loop
		// }

		// var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		// var isValidEmail bool = strings.Contains(email, "@")
		// isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// Boolean variable assignment for city validation
		// isValidCity := city == "Singapore" || city == "London"
		//isInvalidCity := city != "Singapore" && city != "London"
		//isInvalidCity := !isValidCity

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//if userTickets <= remainingTickets {
			//if userTickets < remainingTickets {
			// remainingTickets = remainingTickets - userTickets

			// // bookings[0] = firstName + " " + lastName - array asignment
			// bookings = append(bookings, firstName+" "+lastName) // - slice asignment

			remainingTickets, bookings = bookTicket(userTickets, firstName, lastName, email, bookings)

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Array type: %T\n", bookings)
			// fmt.Printf("Array length: %v\n", len(bookings))

			//Error out of bounds
			//bookings[50] = firstName + " " + lastName

			//userName = "Tom"
			//userTicekts = 2
			fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			var firstNames = getFirstNames(bookings)
			fmt.Printf("The first names on bookings are: %v\n", firstNames)

			// firstNames := []string{}
			// for _, booking := range bookings {
			// 	var names = strings.Fields(booking)
			// 	firstNames = append(firstNames, names[0])
			// }
			//fmt.Printf("These are all our bookings: %v\n", bookings)
			//fmt.Printf("The first names on bookings are: %v\n", firstNames)

			// noTicketsRemaining := remainingTickets == 0 boolean variable assignment
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
			//} else if userTickets == remainingTickets {
			//	remainingTickets = remainingTickets - userTickets

		} else {
			//fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			//break //exit the program
			if !isValidName {
				fmt.Println("Your first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email address is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			// fmt.Println("Your input data is invalid, try again") // General error message

		}
	}

	// Switch statement
	// city := "Singapore"
	// switch city {
	// case "Singapore":
	// 	// Execute this block if city is Singapore
	// case "London":
	// 	// Execute this block if city is London
	// case "Berlin", "Hamburg":
	// 	// Execute this block if city is Berlin or Hamburg
	// default:
	// 	fmt.Println("No valid city selected")
	// }

}

func greetUser(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets to attend!!!")
}

func getFirstNames(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	//fmt.Printf("These are all our bookings: %v\n", bookings)
	//fmt.Printf("The first names on bookings are: %v\n", firstNames)
	return firstNames
}

// func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
// 	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
// 	var isValidEmail bool = strings.Contains(email, "@")
// 	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
// 	return isValidName, isValidEmail, isValidTicketNumber
// }

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	//ask user for number of tickets
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// func bookTicket(userTickets uint, firstName string, lastName string, email string, bookings []string) (uint, []string) {
func bookTicket(userTickets uint, firstName string, lastName string, email string, bookings []map[string]string) (uint, []map[string]string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTicekts"] = strconv.FormatUint(uint64(userTickets), 10)
	//userData["userTickets"] = fmt.Sprint(userTickets)

	// bookings[0] = firstName + " " + lastName - array asignment

	bookings = append(bookings, userData) // - slice asignment
	fmt.Printf("List of bookings: %v\n", bookings)
	//fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive confirmation email at %v.\n", firstName, lastName, userTickets, email)
	//fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets, bookings
}
