package mysql

import (
	"fmt"

	ht "github.com/tarun4all/hotels-golang-app/pkg/hotel"
)

type Storage struct {
	hotels []ht.Hotel
}

func (db *Storage) AddHotel(payload ht.Hotel) error {
	fmt.Println("Add hotel called Payload >> ", payload)
	return nil
}

func (db *Storage) GetHotel() error {
	fmt.Println("Get all called")
	return nil
}
