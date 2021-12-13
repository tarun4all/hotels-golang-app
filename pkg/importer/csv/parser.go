package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadAllRecords(readChannel chan []string) {
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

	// slice data as first row contains column name
	for _, line := range records[1:] {
		readChannel <- line
	}

	fmt.Printf("No. of Scanned records : %v \n", len(records))

	close(readChannel)
}
