package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(base64.StdEncoding.EncodeToString(b))
}
