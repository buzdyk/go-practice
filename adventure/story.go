package main

import (
	"encoding/json"
	"os"
)

type Story map[string]Script

type Script struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"Text"`
	Arc  string `json:"Arc"`
}

func ParseScript() (map[string]Script, error) {
	var r Story
	script, err := os.ReadFile("story.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(script, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
