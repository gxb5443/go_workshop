package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Item struct {
	Title      string
	Link       string
	Date_taken time.Time
}

type Feed struct {
	Title string
	Link  string
	Items []Item
}

func main() {
	resp, err := http.Get("https://api.flickr.com/services/feeds/photos_public.gne?format=json")
	if err != nil {
		log.Fatal("GET FAILED: ", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("HTTP STATUS: ", resp.StatusCode)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal("Print Failed: ", err)
	}
}
