package main

import (
	"fmt"

	gl "github.com/tarun4all/hotels-golang-app/pkg/geolocation"
	rest "github.com/tarun4all/hotels-golang-app/pkg/http/rest"
	storage "github.com/tarun4all/hotels-golang-app/pkg/storage/mysql"
)

func main() {
	storage := storage.New()
	s := gl.NewService(storage)

	// hotel := ht.New()
	// s.AddHotel(hotel)
	fmt.Println(s)

	rest.Init(s)
}
