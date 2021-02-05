package readmoo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HighlightRelationship struct {
	Range struct {
		Data struct {
			Type string
			Id   string
		} `json:"data"`
	} `json:"range"`
}

type HighlightData struct {
	Id           string
	Type         string
	Relationship HighlightRelationship `json:"relationships"`
}

type RangeAttribute struct {
	Content string
}

type BookAttribute struct {
	Title    string
	Subtitle string
	Author   string
	Isbn     string
}

type Include struct {
	Type       string
	Id         string
	Attributes json.RawMessage
}

type Meta struct {
	TotalCount int `json:"total_count"`
}

type HighlightResp struct {
	Meta     Meta
	Data     []HighlightData
	Includes []Include `json:"included"`
}

type ReadingAttribute struct {
	State   string
	Privacy string
}

type ReadingData struct {
	Id           string
	Type         string
	Attribute    ReadingAttribute    `json:"attributes"`
	Relationship ReadingRelationShip `json:"relationships"`
}

type ReadingRelationShip struct {
	Book struct {
		Data struct {
			Type string
			Id   string
		}
	}
}

type ReadingResp struct {
	Meta     Meta
	Data     []ReadingData
	Includes []Include `json:"included"`
}

type Reading struct {
	Id      string
	State   string
	Privacy string
	Book    BookAttribute
}

type Readmoo struct {
	client   http.Client
	apiBase  string
	apiToken string
}

func NewReadmoo(token string) *Readmoo {
	return &Readmoo{
		client:   http.Client{},
		apiBase:  "https://api.readmoo.com/store/v3",
		apiToken: token,
	}
}

func (r *Readmoo) sendRequest(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", "Bearer "+r.apiToken)
	req.Header.Add("Content-Type", "application/vnd.api+json")
	resp, err := r.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Got status %d", resp.StatusCode)
	}
	// TODO handle other status code
	// TODO return error?

	scanner := bufio.NewScanner(resp.Body)
	scanner.Scan()
	body := scanner.Text()

	return body
}

func (r *Readmoo) GetReadingsTotalCount() int {
	url := r.apiBase + "/me/readings/?page[count]=0"

	body := r.sendRequest(url)
	resp := ReadingResp{}
	_ = json.Unmarshal([]byte(body), &resp)

	return resp.Meta.TotalCount
}

func (r *Readmoo) GetReadings() (readings []Reading) {
	totalCount := r.GetReadingsTotalCount()

	apiEntry := r.apiBase + "/me/readings/"
	pageCount := 5
	var url string

	for offset := 0; offset <= totalCount; offset += pageCount {
		url = fmt.Sprintf("%s?page[count]=%d&page[offset]=%d", apiEntry, pageCount, offset)

		body := r.sendRequest(url)
		resp := ReadingResp{}
		_ = json.Unmarshal([]byte(body), &resp)

		for _, datum := range resp.Data {
			var reading Reading
			if datum.Type == "readings" {
				reading.Id = datum.Id
				reading.State = datum.Attribute.State
				reading.Privacy = datum.Attribute.Privacy

				for _, include := range resp.Includes {
					if include.Type == "books" && include.Id == datum.Relationship.Book.Data.Id {
						_ = json.Unmarshal(include.Attributes, &reading.Book)
					}
				}
			}
			readings = append(readings, reading)
		}
	}

	return readings
}

func (r *Readmoo) GetHighlightTotalCount(readingId string) int {
	url := r.apiBase + "/me/readings/" + readingId + "/highlights?page[count]=0"

	body := r.sendRequest(url)
	highlightResp := HighlightResp{}
	_ = json.Unmarshal([]byte(body), &highlightResp)
	return highlightResp.Meta.TotalCount
}

func (r *Readmoo) GetHighlights(readingId string) (highlights []string) {
	totalCount := r.GetHighlightTotalCount(readingId)

	apiEntry := r.apiBase + "/me/readings/" + readingId + "/highlights"
	pageCount := 10
	var url string

	for offset := 0; offset <= totalCount; offset += pageCount {
		url = fmt.Sprintf("%s?page[count]=%d&page[offset]=%d", apiEntry, pageCount, offset)

		body := r.sendRequest(url)
		highlightResp := HighlightResp{}
		_ = json.Unmarshal([]byte(body), &highlightResp)

		for _, highlight := range highlightResp.Data {
			rangeId := highlight.Relationship.Range.Data.Id
			for _, include := range highlightResp.Includes {
				if include.Type == "ranges" && include.Id == rangeId {
					attribute := RangeAttribute{}
					_ = json.Unmarshal(include.Attributes, &attribute)
					highlights = append(highlights, attribute.Content)
				}
			}
		}
	}

	return highlights
}
