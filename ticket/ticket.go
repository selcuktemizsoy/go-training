package ticket

import (
	"example.com/hello/model"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SystemStart(conferenceList *[]model.Conference, userList *[]model.Guest) {

	for checkRemainingTicket(conferenceList) {
		fmt.Println("====Welcome to the Conference Ticket system, we are happy to see you here====")
		fmt.Println("Please select the conference which you want to attend")
		for index, conference := range *conferenceList {
			fmt.Printf("For %v press %v\n", conference.Name, index+1)
		}
		var userChoice, ticketCount int

		fmt.Scan(&userChoice)
		if userChoice < 1 || userChoice > len(*conferenceList) {
			fmt.Println("your input is not correct please enter valid number")
			continue
		}
		selectedConference := &(*conferenceList)[userChoice-1]

		if checkConferenceTicket(selectedConference, 1) {
			fmt.Println("Selected conference ticket is finished, please select different one. Thanks for your " +
				"understanding")
			continue
		}

		for {
			fmt.Println("How many tickets do you want to buy?")
			fmt.Scan(&ticketCount)
			if checkConferenceTicket(selectedConference, ticketCount) {
				fmt.Printf("Available ticket for this conference total: %v please select max this number\n",
					selectedConference.RemainingTicket)
				continue
			} else if ticketCount < 1 {
				fmt.Println("Please select correct number for the ticket count")
			} else {
				break
			}
		}
		var user model.Guest
		fmt.Println("Please enter your name")
		fmt.Scan(&user.FirstName)
		fmt.Println("Please enter your last name")
		fmt.Scan(&user.LastName)

		for {
			fmt.Println("Please enter your email")
			fmt.Scan(&user.Email)
			if !verifyEmailIsValid(&user.Email) {
				fmt.Println("invalid email address. Please try again")
				continue
			}
			break
		}

		user.ReservedTicket = ticketCount
		user.Conference = selectedConference.Name
		selectedConference.RemainingTicket -= ticketCount

		fmt.Println("====Thank you for your interest our conference, we will send you an email with ticket descriptions====")
		fmt.Println("<======================================================================================================>\n\n")
		wg.Add(1)
		go generateTicket(&user, userList, time.Now())
	}
	wg.Wait()
}

func verifyEmailIsValid(email *string) bool {
	return strings.Contains(*email, "@")
}

func checkRemainingTicket(conferenceList *[]model.Conference) bool {
	ticketExist := false
	for _, conference := range *conferenceList {
		if !checkConferenceTicket(&conference, 1) {
			ticketExist = true
			break
		}
	}
	return ticketExist
}

func checkConferenceTicket(conference *model.Conference, requestTicket int) bool {
	return conference.RemainingTicket < requestTicket
}

func generateTicket(user *model.Guest, userList *[]model.Guest, ticketTime time.Time) {
	time.Sleep(3 * time.Second)
	user.TicketNumber = ticketTime.Format("2006040215405000") + strconv.Itoa(rand.Int())
	user.EmailSend = true
	*userList = append(*userList, *user)
	wg.Done()
}
