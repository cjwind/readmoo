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
)

func main() {
	r := readmoo.NewReadmoo("YOUR TOKEN")

	readings := r.GetReadings()
	for _, reading := range readings {
		if reading.State == "finished" {
			highlights := r.GetHighlights(reading.Id)
			fmt.Println(highlights)
		}
	}
}
```

## TODO

- refine error response handling
- convenient way to get bearer token?
