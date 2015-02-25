package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Feed struct {
	Title string
	Link  string
	Items []Item
}

type Item struct {
	Title string
	Link  string
	Tags  string
	Taken time.Time `json:"date_taken"`
}

func main() {
	response, err := http.Get("http://api.flickr.com/services/feeds/photos_public.gne?format=json&tags=hockey&nojsoncallback=1")
	if err != nil {
		log.Fatal("HTTP GET: ", err)
		return
	}
	defer response.Body.Close()
	feed := new(Feed)
	err = json.NewDecoder(response.Body).Decode(feed)
	if err != nil {
		log.Fatal("JSON DECODING: ", err)
		return
	}
	fmt.Println("Feed Title:", feed.Title)
	for _, item := range feed.Items {
		fmt.Println("Title: ", item.Title)
		fmt.Println("Link: ", item.Link)
		fmt.Println("Taken: ", item.Taken)
		fmt.Println()
	}
}
