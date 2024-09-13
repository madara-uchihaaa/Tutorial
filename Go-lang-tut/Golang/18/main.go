package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"plateform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	c1 := Course{
		"a", 1, "a", "a", []string{"a", "a"},
	}
	c2 := Course{
		"b", 2, "b", "b", []string{"b", "b"},
	}

	courses := []Course{c1, c2}

	finalJson, err := json.MarshalIndent(courses, "", "\t")
	// MarshalIndent() is used to format the json data.

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))

}

func DecodeJson() {
	jsonData := []byte(`[
		{
			"coursename": "a",
			"price": 1,
			"plateform": "a",
			"tags": ["a", "a"]
		},
		{
			"coursename": "b",
			"price": 2,
			"plateform": "b",
			"tags": ["b", "b"]
		}
	]`)

	check := json.Valid(jsonData)
	if !check {
		fmt.Println("Invalid json")
		return
	}

	var courses []Course

	err := json.Unmarshal(jsonData, &courses)

	if err != nil {
		panic(err)
	}

	fmt.Println(courses)

	var data map[string]interface{}
	json.Unmarshal([]byte(`{
		"coursename": "a",
		"price": 1,
		"plateform": "a",
		"tags": ["a", "a"]
	}`), &data)

	fmt.Printf("%#v\n", data)
}

func main() {
	// Json in Go;
	// EncodeJson()
	DecodeJson()
}
