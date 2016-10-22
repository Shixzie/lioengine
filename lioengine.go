// Package lioengine is a ml bot that will find updates for the
// project name you give it.
package lioengine

import (
	// "log"
	"strings"
)

// supportedProviders contains all the providers supported by this bot.
var supportedProviders = []string{"Bing", "Twitter"}

func init() {
	// var err error

	// err = fetchKeywords()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// Update is the exported struct that contains
// all kind of info about the project.
type Update struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Link          string   `json:"link"`
	DatePublished string   `json:"date_published"`
	Img           *Img     `json:"img"`
	Category      string   `json:"category"`
	Sources       []string `json:"sources"`

	points int
	words  []byte
	// Maybe more stuff ...
}

// Img contains info about the img/thumbnail
type Img struct {
	Link   string `json:"link"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// replaceSpaces replaces spaces of text with char if the text contains
// spaces.
func replaceSpaces(text, char string) (newText string) {
	// Checks if the text contains spaces
	if strings.Contains(text, " ") {
		// If the text contains spaces it replaces them with char
		newText = strings.Replace(text, " ", char, -1)
		return
	}
	// If text doesn't contain spaces, return the same text
	return text
}
