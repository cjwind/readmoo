# Readmoo

Unofficial API

Able to get:

* Readings
* Higlights of reading

## Usage sample

```go
package main

import (
	"fmt"
	"github.com/cjwind/readmoo"
)

func main() {
	r := readmoo.NewReadmoo("")

	readings := r.GetReadings()
	for _, reading := range readings {
		if reading.State == "finished" {
			highlights := r.GetHighlights(reading.Id)
			fmt.Println(highlights)
		}
	}
}
```
