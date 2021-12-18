package main

import (
	"fmt"
	"log"
	"net/http"

	gl "github.com/tarun4all/hotels-golang-app/pkg/geolocation"
	rest "github.com/tarun4all/hotels-golang-app/pkg/httpHandler/rest"
	storage "github.com/tarun4all/hotels-golang-app/pkg/storage/mysql"
)

func main() {
	storage := storage.New()
	s := gl.NewService(storage)

	// hotel := ht.New()
	// s.AddHotel(hotel)

	handler := rest.NewHandler(s)

	port := ":3001"

	http.Handle("/geodata/", handler)

	fmt.Println("Server starts...")
	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}
