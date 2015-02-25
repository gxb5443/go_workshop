package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Item struct {
	Title string
	Link  string
	Tags  string
	Taken time.Time `json:"date_taken"`
}

type Feed struct {
	Title string
	Link  string
	Items []Item
}

func main() {
	items, err := GetFeed("hockey")
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, item := range items {
		fmt.Println(item.Title)
	}
	items, err = GetFeed("baseball")
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, item := range items {
		fmt.Println(item.Title)
	}
}

func GetFeed(tag string) ([]Item, error) {
	myurl := fmt.Sprintf("http://api.flickr.com/services/feeds/photos_public.gne?format=json&tags=%s&nojsoncallback=1", tag)
	response, err := http.Get(myurl)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, err
	}
	defer response.Body.Close()
	feed := new(Feed)
	err = json.NewDecoder(response.Body).Decode(feed)
	if err != nil {
		return nil, err
	}
	fmt.Println("Feed Title:", feed.Title)
	items := make([]Item, len(feed.Items))
	for i, item := range feed.Items {
		items[i] = item
	}
	return items, nil
}
