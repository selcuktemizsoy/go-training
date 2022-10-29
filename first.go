package main

import (
	"example.com/hello/model"
	"example.com/hello/ticket"
	"example.com/hello/user"
)

func main() {
	goConference := model.Conference{
		Name:            "Go Conference",
		Date:            "27 December",
		TotalTicket:     50,
		RemainingTicket: 50,
	}

	javaConference := model.Conference{
		Name:            "Java Conference",
		Date:            "22 December",
		TotalTicket:     50,
		RemainingTicket: 50,
	}

	dotNetConference := model.Conference{
		Name:            "DotNet Conference",
		Date:            "24 December",
		TotalTicket:     50,
		RemainingTicket: 50,
	}
	conferenceList := make([]model.Conference, 0)
	conferenceList = append(conferenceList, goConference, javaConference, dotNetConference)
	userList := make([]model.Guest, 0)

	ticket.SystemStart(&conferenceList, &userList)
	user.SystemStart(&userList)
}
