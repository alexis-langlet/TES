package main

import (
	holidays "TES/tools"
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload" //load .env file
	"github.com/toqueteos/webbrowser"
)

//go:embed static/*
var static embed.FS

// It takes the start and end dates from the form, loops through each day in between, and if it's not a
// weekend or holiday, it sends a request to the Redmine API to create a time entry for that day
func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")

	startDate, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		log.Fatal(err)
	}

	endDate, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		log.Fatal(err)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Fprintf(w, "time entries created for : \n")
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		stringDate := d.Format("2006-01-02")
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday && !holidays.IsHoliday(stringDate) {
			fmt.Fprintf(w, "%s\n", stringDate)
			sendRequest(stringDate, username, password)

		}
	}

}

// It sends a POST request to the Redmine API with the date, username and password as parameters
func sendRequest(date, username, password string) {

	values := map[string]map[string]interface{}{"time_entry": {"issue_id": 16272, "activity_id": 32, "hours": 7, "spent_on": date}}

	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	redmineAddress := os.Getenv("REDMINE_ADDR")

	resp, err := http.Post("https://"+username+":"+password+"@"+redmineAddress+"/time_entries.json", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
}

func main() {

	fSys, err := fs.Sub(static, "static")

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(fSys)))
	http.HandleFunc("/form", formHandler)

	err = webbrowser.Open("http://localhost:3333")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting server at port 3333\n")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}
}
