// package backend
package main

import (
	bookapi "backend/APP/bookAPI"
	contactapi "backend/APP/contactAPI"
	dashapi "backend/APP/dashAPI"
	loginapi "backend/APP/loginAPI"
	seatsapi "backend/APP/seatsAPI"
	signup "backend/APP/signup"
	ticketapi "backend/APP/ticketAPI"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started........")

	fmt.Println("Connected db........!")
	http.HandleFunc("/register", signup.GetAPI)
	http.HandleFunc("/login", loginapi.GetAPI)
	http.HandleFunc("/dash", dashapi.RetAPI)
	http.HandleFunc("/search", seatsapi.SeatAPI)
	http.HandleFunc("/book", bookapi.BookAPI)
	http.HandleFunc("/ticket", ticketapi.TicketAPI)
	http.HandleFunc("/details", contactapi.DetailAPI)
	http.ListenAndServe(":9091", nil)
}
