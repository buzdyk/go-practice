package main

import (
	json2 "encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestUnmarshalToMap(t *testing.T) {
	json := `{
	"address": {
		"city": "Milan",
		"zip": "00000"
	},
	"name": "John Doe"
}`

	var m map[string]interface{}
	if err := json2.Unmarshal([]byte(json), &m); err != nil {
		fmt.Println("Unable to decode json:", err.Error())
		os.Exit(1)
	}

	t.Log("Name:", m["name"])
	t.Log("City:", m["address"].(map[string]interface{})["city"])
	t.Log("Zip:", m["address"].(map[string]interface{})["zip"])
}
