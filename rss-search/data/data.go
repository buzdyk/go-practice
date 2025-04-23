package data

import (
	"encoding/json"
	"os"
)

type Feed struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func LoadData() ([]*Feed, error) {
	file, err := os.Open("data/data.json")

	if err != nil {
		return nil, err
	}

	var feeds []*Feed

	err = json.NewDecoder(file).Decode(&feeds)

	if err != nil {
		return nil, err
	}

	return feeds, err
}
