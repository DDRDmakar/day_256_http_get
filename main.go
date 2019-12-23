package main

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"encoding/json"
)

// HTTP request handler
func requestHandler(w http.ResponseWriter, r *http.Request) {
	
	type responseStruct struct {
		ErrorCode   int    `json:"errorCode"`
		DataMessage string `json:"dataMessage"`
	}
	
	Ystr := r.URL.Query().Get("year") // Year we got from request
	
	const dayNum int = 256
	Y, err := strconv.Atoi(Ystr) // Get year as int
	
	if err == nil {
		// Можно бы было в аргументе "дни" просто указать 256, но мы лёгких путей не ищем!
		t := time.Date(Y, time.January, 0, 0, 0, 0, 0, time.UTC)
		// Duration - 256 days
		// 86400 seconds in one day
		// 86400 * 256 = 22118400 seconds
		dur, _ := time.ParseDuration("22118400s")
		t := t.Add(dur) // Add 256 days to the beginning of our year
		timestr := t.Format("02/01/06") // Convert it to string DD/MM/YY
		
		answer := responseStruct {
			ErrorCode: 200,
			DataMessage: timestr,
		}
		err := json.NewEncoder(w).Encode(&answer) // Send answer back
		if err != nil {
			fmt.Println("Error sending JSON")
		}
	}
	else {
		fmt.Println("Error parsing year integer")
	}
}

func main() {
	// This server spawns asynchronous goroutine
	// for each HTTP request
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":80", nil) // Here we wait and listen
	fmt.Println(err)
}
