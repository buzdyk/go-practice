// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	redirects()
}

func readBody() {
	res, err := http.Get("https://practicum.yandex.ru")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	io.CopyN(os.Stdout, res.Body, 512)

	//b := make([]byte, 512)
	//res.Body.Read(b)
	//fmt.Println(b, string(b))
}

func redirects() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL)
			return nil
		},
	}

	client.Get("http://ya.ru")
}
