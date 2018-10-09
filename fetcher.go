package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Category struct {
	Term string `xml:"term,attr"`
}
type Entry struct {
	Title    string   `xml:"title"`
	Category Category `xml:"category"`
}
type Feed struct {
	Name    string  `xml:"title"`
	Updated string  `xml:"updated"`
	Entries []Entry `xml:"entry"`
}

func main() {
	vancouverUrl := "https://weather.gc.ca/rss/city/bc-74_e.xml"
	response, err := http.Get(vancouverUrl)
	if err != nil {
		fmt.Printf("error getting weather data: %T: %+v\n", err, err)
		os.Exit(1)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error reading weather response: %T: %+v\n", err, err)
		os.Exit(1)
	}

	var feed Feed
	if err := xml.Unmarshal(contents, &feed); err != nil {
		fmt.Printf("error processing XML response: %T: %+v\n", err, err)
		os.Exit(1)
	}

	currentConditions := feed.findCurrentConditions()
	fmt.Println(currentConditions.Title)
}

func (f *Feed) findCurrentConditions() *Entry {
	for _, entry := range f.Entries {
		fmt.Printf("entry: %T: %+v\n", entry, entry)
		if entry.Category.Term == "Current Conditions" {
			return &entry
		}
	}
	return nil
}
