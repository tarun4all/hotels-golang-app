package main

import (
	"fmt"
	"time"
)

type Hotel struct {
	IpAddress    string
	CountryCode  string
	Country      string
	City         string
	Latitude     float64
	Longitude    float64
	MysteryValue int64
	Created      time.Time
}

type Storage struct {
	hotels []Hotel
}

type Service interface {
	AddHotel(Hotel) error
	AddSampleHotels([]Hotel)
}

func (s *service) AddHotel(payload Hotel) error {
	fmt.Println("Add hotel service called", payload)

	return nil
}

// AddSampleHotels adds some sample Hotels to the database
func (s *service) AddSampleHotels(payload []Hotel) {
	fmt.Println("Add Sample hotel service called")
}

type Repository interface {
	addHotel(Hotel) error
	getHotel() error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (db Storage) addHotel(payload Hotel) error {
	fmt.Println("Add hotel called Payload >> ", payload)
	return nil
}

func (db Storage) getHotel() error {
	fmt.Println("Get all called")
	return nil
}

func main() {
	storage := new(Storage)
	s := NewService(storage)
	s.AddHotel(Hotel{
		IpAddress:    "192.168.0.1",
		CountryCode:  "IN",
		Country:      "INdia",
		City:         "Delhi",
		Latitude:     78.888888,
		Longitude:    78.888888,
		MysteryValue: 123456,
		Created:      time.Now(),
	})
}
