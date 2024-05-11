package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Reading struct {
	Time        time.Time
	Temperature float64
	Humidity    float64
	Source      string
}

var db *sql.DB
var err error

func init() {
	connStr := "postgres://postgres:mysecretpassword@localhost/postgres?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func GetLatest(source string) (Reading, error) {
	rows, err := db.Query(`
		SELECT * FROM readings
		WHERE source = $1
		ORDER BY DATETIME DESC LIMIT 1
	`, source)
	if err != nil {
		return Reading{}, err
	}
	defer rows.Close()

	out := Reading{}
	if rows.Next() {
		err = rows.Scan(&out.Time, &out.Temperature, &out.Humidity, &out.Source)
		if err != nil {
			return out, err
		}
	}

	return out, nil
}

func AddReading(val Reading) error {
	fmt.Println(val)
	// TODO: call sql query to insert the value as a row
	// return error if query fails
	return nil
}
