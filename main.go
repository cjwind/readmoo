package main

import (
	"fmt"
	readmoo2 "github.com/cjwind/readmoo-api/readmoo"
)

func main() {
	readmoo := readmoo2.Readmoo{
		ApiBase:  "https://api.readmoo.com/store/v3",
		ApiToken: "",
	}

	readings := readmoo.GetReadings()
	for _, reading := range readings {
		if reading.State == "finished" {
			highlights := readmoo.GetHighlights(reading.Id)
			fmt.Println(highlights)
		}
	}
}
