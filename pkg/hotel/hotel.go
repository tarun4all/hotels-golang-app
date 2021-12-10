package hotel

import (
	"strconv"
)

type Hotel struct {
	IpAddress   string  `json:"ipAddress"`
	CountryCode string  `json:"countryCode"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Created     string  `json:"createdAt"`
}

func New(ipAddress, countryCode, country, city, latitude, longitude, createdAt string) Hotel {
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		lat = 0
	}

	long, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		long = 0
	}

	return Hotel{
		IpAddress:   ipAddress,
		CountryCode: countryCode,
		Country:     country,
		City:        city,
		Latitude:    lat,
		Longitude:   long,
		Created:     createdAt,
	}
}

type HotelStorage interface {
	AddHotel(Hotel) error
	GetHotel(string) ([]Hotel, error)
}

type HotelService struct {
	storage HotelStorage
}

func (s *HotelService) AddHotel(payload Hotel) error {
	return s.storage.AddHotel(payload)
}

// AddSampleHotels adds some sample Hotels to the database
func (s *HotelService) GetHotel(requestIPAddr string) ([]Hotel, error) {
	info, err := s.storage.GetHotel(requestIPAddr)

	if err != nil {
		panic(err.Error())
	}
	return info, nil
}

func NewService(st HotelStorage) *HotelService {
	return &HotelService{st}
}
