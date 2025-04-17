package main

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	data := url.Values{}
	data.Set("key1", "value1")
	data.Set("key2", "value2")

	request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))

	client := http.Client{}
	client.Do(request)
}
