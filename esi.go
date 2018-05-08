package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"time"
	"sort"
	"strings"
)

var (
	Home int32 = 31002443
)

type Bookmark struct {
	BookmarkID int32 `json:"bookmark_id"`
	Coordinates struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"coordinates"`
	Created    Time   `json:"created"`
	CreatorID  int32  `json:"creator_id"`
	Label      string `json:"label"`
	LocationID int32  `json:"location_id"`
	Notes      string `json:"notes"`
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var rawString string
	if err := json.Unmarshal(data, &rawString); err != nil {
		return err
	}

	tt, err := time.Parse(time.RFC3339, rawString)
	if err != nil {
		return err
	}

	t.Time = tt

	return nil
}

func TodayBookmark(b []Bookmark) []Bookmark {
	if len(b) == 0 {
		return b
	}

	now := time.Now()
	rounded := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, b[0].Created.Location())
	endId := len(b)

	for i, bookmark := range b {
		if bookmark.Created.Before(rounded) {
			endId = i
			break
		}
	}

	return b[:endId]
}

func IsStaticWH(label string) bool {
	label = strings.ToLower(label)



	return false
}

func DetectedStaticWH(b []Bookmark) []Bookmark {
	firstHomeID := 0

	for i, bookmark := range b {
		if bookmark.LocationID == Home {
			firstHomeID = i
			//break


		}
	}

	return b[:firstHomeID]
}

func main() {
	dat, _ := ioutil.ReadFile("response.txt")

	var decoded []Bookmark

	if err := json.Unmarshal(dat, &decoded); err != nil {
		panic(err)
	}

	sort.Slice(decoded, func(i, j int) bool { return decoded[i].Created.Time.After(decoded[j].Created.Time) })

	decoded = TodayBookmark(decoded)

	for _, bookmark := range decoded {
		if bookmark.LocationID == Home {
			fmt.Println("Home", bookmark.Label, bookmark.Created.Format(time.RFC850))
		} else {
			fmt.Println(bookmark.LocationID, bookmark.Label, bookmark.Created.Format(time.RFC850))
		}
	}
}
