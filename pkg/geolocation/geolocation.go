package geolocation

import (
	"strconv"
)

type Geolocation struct {
	IpAddress   string  `json:"ipAddress"`
	CountryCode string  `json:"countryCode"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Created     string  `json:"createdAt"`
}

func New(ipAddress, countryCode, country, city, latitude, longitude, createdAt string) Geolocation {
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		lat = 0
	}

	long, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		long = 0
	}

	return Geolocation{
		IpAddress:   ipAddress,
		CountryCode: countryCode,
		Country:     country,
		City:        city,
		Latitude:    lat,
		Longitude:   long,
		Created:     createdAt,
	}
}

type GeolocationStorage interface {
	AddGeolocation(Geolocation) error
	GetGeolocation(string) ([]Geolocation, error)
}

type GeolocationService struct {
	storage GeolocationStorage
}

func (s *GeolocationService) AddGeolocation(payload Geolocation) error {
	return s.storage.AddGeolocation(payload)
}

// AddSampleGeolocations adds some sample Geolocations to the database
func (s *GeolocationService) GetGeolocation(requestIPAddr string) ([]Geolocation, error) {
	info, err := s.storage.GetGeolocation(requestIPAddr)

	if err != nil {
		panic(err.Error())
	}
	return info, nil
}

func NewService(st GeolocationStorage) *GeolocationService {
	return &GeolocationService{st}
}
