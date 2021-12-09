package memory

import (
	"fmt"

	ht "github.com/tarun4all/hotels-golang-app/pkg/hotel"
)

type Storage struct {
	hotels []ht.Hotel
}

func (db *Storage) AddHotel(payload ht.Hotel) error {
	fmt.Println("Add hotel Memory called >> ", payload)
	return nil
}

func (db *Storage) GetHotel() error {
	fmt.Println("Get all called")
	return nil
}

func New() *Storage {
	return &Storage{}
}
