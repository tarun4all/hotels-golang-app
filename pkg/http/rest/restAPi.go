package rest

import (
	"fmt"
	"log"
	"net/http"

	ht "github.com/tarun4all/hotels-golang-app/pkg/hotel"
)

func GetHotels(hotelService *ht.HotelService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := hotelService.GetHotel()
		fmt.Printf("%+v", r)
		fmt.Println(">>", err)
		fmt.Fprintf(w, "Hotels List")
	}
}

func Init(hotelService *ht.HotelService) {
	fmt.Println("Setting up router...")
	http.HandleFunc("/hotels", GetHotels(hotelService))
	log.Fatal(http.ListenAndServe(":3001", nil))
}
