package main

import (
	"fmt"
)

func main() {
	var readmoo Readmoo
	readmoo.ApiBase = "https://api.readmoo.com/store/v3"
	readmoo.ApiToken = ""

	readings := readmoo.getReadings()
	for _, reading := range readings {
		if reading.State == "finished" {
			highlights := readmoo.getHighlights(reading.Id)
			fmt.Println(highlights)
		}
	}
}
