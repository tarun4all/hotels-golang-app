package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	gl "github.com/tarun4all/hotels-golang-app/pkg/geolocation"
	importer "github.com/tarun4all/hotels-golang-app/pkg/importers"
	storage "github.com/tarun4all/hotels-golang-app/pkg/storage/mysql"
)

func main() {
	DB_URL := os.Getenv("DB_URL")
	pathArg := os.Args[1:]

	fmt.Println("Connecting db ... ", DB_URL)
	storage := storage.New(DB_URL)
	s := gl.NewService(storage)
	csvImporter := importer.NewCsvImporter()

	if len(pathArg) == 0 {
		log.Fatal("No csv path provided")
	}

	readChannel, err := csvImporter.Import(pathArg[0])

	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Wait()

	for data := range readChannel {
		defer wg.Done()
		wg.Add(1)
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
