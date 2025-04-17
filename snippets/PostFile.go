package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	filename := "PostFile.go"
	file, _ := os.Open(filename)
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("uploadfile", filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	writer.Close()

	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	client.Do(request)
}
