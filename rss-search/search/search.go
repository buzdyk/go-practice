package search

import (
	"fmt"
	"github.com/buzdyk/rss-search/data"
)

type Result struct {
	Url         string
	Description string
}

func Run(search string) {
	feeds, err := data.LoadData()

	if err != nil {
		panic(err)
	}

	var results chan []*Result

	for _, feed := range feeds {

		fmt.Println(feed.URL, feed.Title)
	}
}
