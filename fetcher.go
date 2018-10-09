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
		printErrAndExit(err, "error getting weather data")
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		printErrAndExit(err, "error reading weather response")
	}

	var feed Feed
	if err := xml.Unmarshal(contents, &feed); err != nil {
		printErrAndExit(err, "error processing XML response")
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

func printErrAndExit(e error, message string) {
	fmt.Printf("%v: %T: %+v\n", message, e, e)
	os.Exit(1)
}
