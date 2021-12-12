package main

import (
	gl "github.com/tarun4all/hotels-golang-app/pkg/geolocation"
	parser "github.com/tarun4all/hotels-golang-app/pkg/importer/csv"
	storage "github.com/tarun4all/hotels-golang-app/pkg/storage/mysql"
)

func main() {
	storage := storage.New()
	s := gl.NewService(storage)

	parser.ParseCSV(s)
}
