package search

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/buzdyk/rss-search/data"
	"net/http"
)

type (
	// item defines the fields associated with the item tag
	// in the rss document.
	item struct {
		PubDate     string `xml:"pubDate"`
		Description string `xml:"description"`
		Link        string `xml:"link"`
		GUID        string `xml:"guid"`
	}

	channel struct {
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        string `xml:"link"`
		Item        []item `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document.
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

func SearchFeed(feed *Feed, search string, results chan []*Result) {

}

func retrieve(feed *data.Feed) (*rssDocument, error) {
	if feed.URL == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	resp, err := http.Get(feed.URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
