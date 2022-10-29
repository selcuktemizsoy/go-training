package model

type Conference struct {
	Name            string
	Date            string
	TotalTicket     int
	RemainingTicket int
}

type Guest struct {
	FirstName, LastName, Email string
	ReservedTicket             int
	TicketNumber               string
	EmailSend                  bool
	Conference                 string
}
