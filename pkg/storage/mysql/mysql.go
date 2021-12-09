package mysql

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	ht "github.com/tarun4all/hotels-golang-app/pkg/hotel"
)

type Storage struct {
	db *sql.DB
}

func (db *Storage) AddHotel(payload ht.Hotel) error {
	fmt.Println("Add hotel called Payload >> ", payload)

	insertQuery, _, _ := sq.Insert("hotels").Columns("name", "age").Values("moe", 13).ToSql()
	fmt.Println(insertQuery)

	insert, err := db.db.Query(insertQuery, "Tarun", 10)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	return nil
}

func (db *Storage) GetHotel() error {
	query := sq.Select("*").From("hotels")
	query = query.Where(sq.Eq{"id": nil})
	sql, _, _ := query.ToSql()
	fmt.Println(sql, *db)

	return nil
}

func New() *Storage {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	return &Storage{db}
}
