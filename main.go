package main

import (
	"fmt"
	"github.com/cjwind/readmoo-api/readmoo"
)

func main() {
	r := readmoo.Readmoo{
		ApiBase:  "https://api.r.com/store/v3",
		ApiToken: "",
	}

	readings := r.GetReadings()
	for _, reading := range readings {
		if reading.State == "finished" {
			highlights := r.GetHighlights(reading.Id)
			fmt.Println(highlights)
		}
	}
}
