package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	gl "github.com/tarun4all/hotels-golang-app/pkg/geolocation"
)

func ParseCSV(geolocationService *gl.GeolocationService) {
	csvFile, err := os.Open("./data_dump.csv")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Processing CSV file...")
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("error ", err)
	}

	var validRecord = 0

	// slice data as first row contains column name
	for _, line := range records[1:] {
		geolocation := gl.Geolocation{}
		err := geolocation.Parse(line)

		if err != nil {
			continue
		}

		err = geolocation.ValidatePayload()

		if err != nil {
			continue
		}

		err = geolocationService.AddGeolocation(geolocation)

		if err != nil {
			continue
		}

		validRecord++
	}

	fmt.Printf("No. of Scanned records : %v \n", len(records))
	fmt.Printf("No. of Valid records : %v \n", validRecord)
	fmt.Printf("No. of Invalid record : %v \n", len(records)-validRecord)
}
