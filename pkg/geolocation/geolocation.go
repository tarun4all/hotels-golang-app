package geolocation

import (
	"errors"
	"net"
	"regexp"
	"strconv"
	"time"
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

func (geolocation *Geolocation) Parse(data []string) error {
	lat, err := strconv.ParseFloat(data[4], 64)
	if err != nil {
		return errors.New("Invalid Latitude")
	}

	long, err := strconv.ParseFloat(data[5], 64)
	if err != nil {
		return errors.New("Invalid Longitude")
	}

	geolocation.IpAddress = data[0]
	geolocation.CountryCode = data[1]
	geolocation.Country = data[2]
	geolocation.City = data[3]
	geolocation.Latitude = lat
	geolocation.Longitude = long
	geolocation.Created = time.Now().String()

	return nil
}

func (geolocation *Geolocation) ValidatePayload() error {
	if net.ParseIP(geolocation.IpAddress) == nil {
		return errors.New("Invalid IP")
	}

	if ok, err := regexp.MatchString("\\b[A-Z]{2,2}$\\b", geolocation.CountryCode); ok == false || err != nil {
		return errors.New("Invalid Country Code")
	}

	if geolocation.Country == "" {
		return errors.New("Invalid Country")
	}

	if geolocation.City == "" {
		return errors.New("Invalid City")
	}

	// if string
	// if _, err := strconv.ParseFloat(geolocation.Latitude, 64); err != nil {
	// 	return errors.New("Invalid Latitude")
	// }

	// if _, err := strconv.ParseFloat(payload[5], 64); err != nil {
	// 	return errors.New("Invalid Longitude")
	// }

	if geolocation.Longitude > 90 && geolocation.Longitude < -90 {
		return errors.New("Invalid Latitude")
	}

	if geolocation.Longitude > 180 && geolocation.Longitude < -180 {
		return errors.New("Invalid Longitude")
	}

	return nil
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
