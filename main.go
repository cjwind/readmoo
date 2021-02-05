package main

import (
	"fmt"
)

func main() {
	readmoo := Readmoo{
		ApiBase:  "https://api.readmoo.com/store/v3",
		ApiToken: "",
	}

	readings := readmoo.getReadings()
	for _, reading := range readings {
		if reading.State == "finished" {
			highlights := readmoo.getHighlights(reading.Id)
			fmt.Println(highlights)
		}
	}
}
