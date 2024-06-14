package main

import "testing"

func TestMain(t *testing.T) {

	main()
}

func TestBook(t *testing.T) {

	ebookMiddle := EBookMiddle{
		Meta: Meta{
			Authors: []struct {
				Name string `json:"name"`
			}{
				{
					Name: "Ewy",
				},
			},
			Tags: []struct {
				Name string `json:"name"`
			}{
				{
					Name: "horror",
				},
			},
			Title:    "my butthole",
			Subtitle: "sdoifjodsij",
			Date:     "2024-06-14",
			Language: "English",
			Rights:   "copyright 2024",
		},
	}

	createHTMLBook(ebookMiddle)
}
