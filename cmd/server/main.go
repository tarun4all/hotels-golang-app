package main

import (
	ht "github.com/tarun4all/hotels-golang-app/pkg/hotel"
	storage "github.com/tarun4all/hotels-golang-app/pkg/storage/memory"
)

func main() {
	storage := storage.New()
	s := ht.NewService(storage)

	hotel := ht.New()
	s.AddHotel(hotel)
}
