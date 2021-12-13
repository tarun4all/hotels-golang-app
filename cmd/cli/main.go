package main

import (
	"fmt"

	gl "github.com/tarun4all/hotels-golang-app/pkg/geolocation"
	parser "github.com/tarun4all/hotels-golang-app/pkg/importer/csv"
	storage "github.com/tarun4all/hotels-golang-app/pkg/storage/mysql"
)

func main() {
	storage := storage.New()
	s := gl.NewService(storage)

	// channel for csv rows
	readChannel := make(chan []string)

	go func() {
		var validRecord = 0
		for data := range readChannel {
			geolocation := gl.Geolocation{}
			err := geolocation.Parse(data)

			if err != nil {
				continue
			}

			err = geolocation.ValidatePayload()

			if err != nil {
				continue
			}

			err = s.AddGeolocation(geolocation)

			if err != nil {
				continue
			}

			validRecord++
		}

		fmt.Printf("No. of Valid records : %v \n", validRecord)
	}()

	parser.ReadAllRecords(readChannel)
}
