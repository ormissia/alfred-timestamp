package main

import (
	"encoding/json"
	"fmt"
)

type Output struct {
	Items []Item `json:"items"`
}

type Item struct {
	Uid          string `json:"uid"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Arg          string `json:"arg"`
	Autocomplete string `json:"autocomplete"`
	Icon         struct {
		Type string `json:"type"`
		Path string `json:"path"`
	} `json:"icon"`
}

func OutPut(outString []string) {
	output := Output{Items: make([]Item, 0)}

	for _, s := range outString {
		output.Items = append(output.Items, Item{
			Type:         s,
			Title:        s,
			Subtitle:     s,
			Arg:          s,
			Autocomplete: s,
		})
	}

	outputBytes, err := json.Marshal(output)
	if err != nil {
		return
	}
	fmt.Println(string(outputBytes))
}
