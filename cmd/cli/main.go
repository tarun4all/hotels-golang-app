package main

import (
	"fmt"
	"os"

	gl "github.com/tarun4all/hotels-golang-app/pkg/geolocation"
	importer "github.com/tarun4all/hotels-golang-app/pkg/importers"
	storage "github.com/tarun4all/hotels-golang-app/pkg/storage/mysql"
)

func main() {
	DB_URL := os.Getenv("DB_URL")

	storage := storage.New(DB_URL)
	s := gl.NewService(storage)
	csvImporter := importer.NewCsvImporter()

	readChannel, err := csvImporter.Import("../../data_dump.csv")

	if err != nil {
		fmt.Println(err)
	}

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

	}
}
