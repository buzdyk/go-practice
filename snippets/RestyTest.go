package main

import (
	"fmt"
	"resty.dev/v3"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	client := resty.New()
	defer client.Close()

	var users []User

	res, err := client.R().
		EnableTrace().
		SetResult(&users).
		Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		panic(err)
	}

	fmt.Println(users)

	defer res.Body.Close()
}
