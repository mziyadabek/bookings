package main

import (
	"github.com/mziyadabek/bookings/pkg/handlers"
	"net/http"
)

const PortNum = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	_ = http.ListenAndServe(PortNum, nil)
}
