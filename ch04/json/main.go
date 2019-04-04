package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	TotalCount int
	Movies     []*Movie `json:"movies"`
}

type Movie struct {
	Title string `json:"title"`
}

var response string = `
	{
		"TotalCount": 2,
		"movies": [
			{ "title": "The Martian" },
			{ "title": "SOCIAL NETWORK" }
		]
	}
`

func main() {
	var result Result
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		fmt.Println(err)
	}
	for _, movie := range result.Movies {
		fmt.Println(*movie)
	}
}
