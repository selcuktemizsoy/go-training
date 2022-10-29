package user

import (
	"example.com/hello/model"
	"fmt"
)

func SystemStart(userList *[]model.Guest) {
	fmt.Println("Here is the list of tickets and user information for our conferences")
	for index, user := range *userList {
		fmt.Printf("User number %v\n", index+1)
		var isEmailSend string
		if user.EmailSend {
			isEmailSend = "yes"
		} else {
			isEmailSend = "no"
		}
		fmt.Printf("Name: %v, LastName: %v, ticket count: %v, ticker number number: %v is email sent: %v\n", user.FirstName, user.LastName,
			user.ReservedTicket, user.TicketNumber, isEmailSend)
	}
}
