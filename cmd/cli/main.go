package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Hotel struct {
	IpAddress    string
	CountryCode  string
	Country      string
	City         string
	Latitude     float64
	Longitude    float64
	MysteryValue int64
}

func dataFromatter(data []string) Hotel {
	lat, err := strconv.ParseFloat(data[4], 64)
	if err != nil {
		lat = 0
	}

	long, err := strconv.ParseFloat(data[5], 64)
	if err != nil {
		long = 0
	}

	mv, err := strconv.ParseInt(data[6], 0, 64)
	if err != nil {
		mv = 0
	}

	return Hotel{
		IpAddress:    data[0],
		CountryCode:  data[1],
		Country:      data[2],
		City:         data[3],
		Latitude:     lat,
		Longitude:    long,
		MysteryValue: mv,
	}
}

func parseCSV(filePath string, formatter func([]string) Hotel) {
	csvFile, err := os.Open(filePath)
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
	fmt.Println(records[0])
	for _, line := range records {

		fmt.Printf("%+v\n", formatter(line))
	}
}

func main() {
	parseCSV("./data_dump.csv", dataFromatter)
}
