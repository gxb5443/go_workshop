package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Photo struct {
	Title      string
	Link       string
	Tags       string
	Date_taken time.Time
	Author     string
}

type Feed struct {
	Title string  `json:"title"`
	Link  string  `json:"link"`
	Items []Photo `json:"items"`
}

func main() {
	resp, err := http.Get("http://api.flickr.com/services/feeds/photos_public.gne?format=json&tags=nyc&nojsoncallback=1")
	if err != nil {
		log.Fatal("GET FAILED: ", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("HTTP STATUS: ", resp.StatusCode)
	}
	feed := new(Feed)
	err = json.NewDecoder(resp.Body).Decode(feed)
	if err != nil {
		log.Fatal("JSON FAILED: ", err)
	}
	fmt.Println(feed.Title)
	for _, child := range feed.Items {
		fmt.Println("Title: ", child.Title)
		fmt.Println("Author: ", child.Author)
		fmt.Println("Tags: ", child.Tags)
		fmt.Println()
	}
}
