# Readmoo

Unofficial API

Able to get:

* Readings
* Higlights of reading

You need to get your bearer token first.

## Usage sample

```go
package main

import (
	"fmt"
	"github.com/cjwind/readmoo"
	"log"
)

func main() {
	r := readmoo.NewReadmoo("YOUR TOKEN")

	readings, err := r.GetReadings()

	if err != nil {
		log.Fatalln(err)
	}

	for _, reading := range readings {
		if reading.State == "finished" {
			highlights, err := r.GetHighlights(reading.Id)

			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(highlights)
		}
	}
}
```

## TODO

- convenient way to get bearer token?
