package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type ReadingRow struct {
	Time        time.Time
	Location    string
	Temperature float64
	Humidity    float64
}

func main() {
	// connect to the database running on localhost using password
	// run query SELECT * from readings;
	// print query results
	connStr := "postgres://postgres:mysecretpassword@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * from readings where location = 'fridge' limit 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var reading ReadingRow
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(&reading.Time, &reading.Location, &reading.Temperature, &reading.Humidity); err != nil {
			log.Fatal(err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(reading)
}
