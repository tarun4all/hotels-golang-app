package main

import (
	"fmt"

	ht "github.com/tarun4all/hotels-golang-app/pkg/hotel"
	rest "github.com/tarun4all/hotels-golang-app/pkg/http/rest"
	storage "github.com/tarun4all/hotels-golang-app/pkg/storage/mysql"
)

func main() {
	storage := storage.New()
	s := ht.NewService(storage)

	// hotel := ht.New()
	// s.AddHotel(hotel)
	fmt.Println(s)

	rest.Init(s)
}
