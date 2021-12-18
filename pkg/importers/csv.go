package importers

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type CSVImporter struct {
	metrics Metrics
}

func (c *CSVImporter) GetMetrics() Metrics {
	return c.metrics
}

func NewCsvImporter() Importer {
	importer := new(CSVImporter)
	return importer
}

func (importer *CSVImporter) Import(path string) (<-chan []string, error) {
	csvFile, err := os.Open(path)

	// channel for csv rows
	readChannel := make(chan []string)

	if err != nil {
		return readChannel, err
	}

	fmt.Println("Processing CSV file...")

	csvReader := csv.NewReader(csvFile)
	header, err := csvReader.Read()

	fmt.Println("Header >> ", header)
	if err != nil {
		return readChannel, err
	}

	go func() {
		defer csvFile.Close()
		defer close(readChannel)

		for {
			record, err := csvReader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println("error ", err)
			}

			readChannel <- record
		}
	}()

	return readChannel, nil
}
