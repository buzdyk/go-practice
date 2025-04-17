package main

import (
	"bytes"
	"net/http"
)

func main() {
	url := "https://google.com"

	var body = []byte(`{"message":"Hello"}`)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	client.Do(request)
}
