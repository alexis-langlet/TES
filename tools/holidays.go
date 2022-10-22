package holidays

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getHolydays() []string {
	response, err := http.Get(fmt.Sprintf("https://calendrier.api.gouv.fr/jours-feries/metropole/%d.json", time.Now().Year()))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Declared an empty map interface
	var responseData map[string]interface{}

	r, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(r, &responseData)

	i := 0
	holidays := make([]string, len(responseData))
	for k := range responseData {
		holidays[i] = k
		i++
	}

	return holidays
}

func IsHoliday(date string) bool {
	for _, holiday := range getHolydays() {
		if date == holiday {
			return true
		}
	}
	return false
}
