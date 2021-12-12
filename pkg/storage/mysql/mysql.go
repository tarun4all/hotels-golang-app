package mysql

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	gl "github.com/tarun4all/hotels-golang-app/pkg/geolocation"
)

type Storage struct {
	db *sql.DB
}

func (db *Storage) AddGeolocation(payload gl.Geolocation) error {
	insertQuery, args, _ := sq.Insert("geolocation").Columns("ipAddress", "countryCode", "country", "city", "latitude", "longitude", "createdAt").Values(payload.IpAddress, payload.CountryCode, payload.Country, payload.City, payload.Latitude, payload.Longitude, payload.Created).ToSql()
	fmt.Println(insertQuery)

	insert, err := db.db.Query(insertQuery, args...)
	fmt.Println(" >>> ", insert)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	return nil
}

func (db *Storage) GetGeolocation(requestIPAddr string) ([]gl.Geolocation, error) {
	query := sq.Select("*").From("geolocation")
	query = query.Where(sq.Eq{"ipAddress": requestIPAddr})
	sql, args, _ := query.ToSql()

	rows, err := db.db.Query(sql, args...)

	var geolocations []gl.Geolocation
	defer rows.Close()
	for rows.Next() {
		var ipAddress, countryCode, country, city, latitude, longitude, createdAt string
		if err := rows.Scan(&ipAddress, &countryCode, &country, &city, &latitude, &longitude, &createdAt); err != nil {
			return geolocations, err
		}
		geolocations = append(geolocations, gl.New(ipAddress, countryCode, country, city, latitude, longitude, createdAt))
	}

	if err != nil {
		return geolocations, err
	}
	return geolocations, nil
}

func New() *Storage {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	return &Storage{db}
}
