package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	title  string `json:"title"`
	author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {

	author := Author{
		Name:      "joseph",
		Age:       22,
		Developer: true,
	}

	book := Book{
		title:  "learning go ",
		author: author,
	}

	fmt.Printf("%+v\n", book)

	byteArr, err := json.MarshalIndent(book, "", "   ")

	if err != nil {
		fmt.Println(err, "fall in error")
	}

	fmt.Println(string(byteArr))
}
