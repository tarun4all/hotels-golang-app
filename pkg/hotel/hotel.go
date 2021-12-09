package hotel

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

func New() Hotel {
	return Hotel{
		IpAddress:    "192.168.0.1",
		CountryCode:  "IN",
		Country:      "INdia",
		City:         "Delhi",
		Latitude:     78.888888,
		Longitude:    78.888888,
		MysteryValue: 123456,
		Created:      time.Now(),
	}
}

type HotelStorage interface {
	AddHotel(Hotel) error
	GetHotel() error
}

type HotelService struct {
	storage HotelStorage
}

func (s *HotelService) AddHotel(payload Hotel) error {
	fmt.Println("Add hotel service called", payload)
	return s.storage.AddHotel(payload)
}

// AddSampleHotels adds some sample Hotels to the database
func (s *HotelService) GetHotel() error {
	fmt.Println("Add Sample hotel service called")
	return nil
}

func NewService(st HotelStorage) *HotelService {
	return &HotelService{st}
}
