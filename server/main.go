package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
	  <html>
	     <title>Temp mon</title>
         <body>
            <h1>Temp Mon</h1>
			<p>TODO</p>
		 </body>
	  </html>
	`)
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "ok"}`)
}

func postReading(w http.ResponseWriter, req *http.Request) {
	// TODO: parse request body and insert into db
	reading := Reading{
		Time:        time.Time{},
		Temperature: 0,
		Humidity:    0,
		Source:      "",
	}
	err := AddReading(reading)
	if err != nil {
		log.Println(err)
		return
	}
}

func getLatest(w http.ResponseWriter, req *http.Request) {
	// read request query param for source value
	source := req.URL.Query().Get("source")
	reading, err := GetLatest(source)
	if err != nil {
		log.Println(err)
		return
	}

	marhsalled, err := json.Marshal(reading)
	if err != nil {
		log.Println(err)
		return
	}

	// TODO: select from db instead of hardcoded value
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(marhsalled))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health", health)
	http.HandleFunc("/readings/add", postReading)
	http.HandleFunc("/readings/latest", getLatest)
	http.ListenAndServe(":8090", nil)
}
