package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type RSSEnclosure struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`

	Title         string `xml:"channel>title"`
	Link          string `xml:"channel>link"`
	Description   string `xml:"channel>description"`
	LastBuildDate string `xml:"channel>lastBuildDate"`
	PubDate       string `xml:"channel>pubDate"`
	ItemList      []Item `xml:"channel>item"`
}

type Item struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	Content     template.HTML `xml:"encoded"`
	PubDate     string        `xml:"pubDate"`
	Comments    string        `xml:"comments"`
}

func main() {
	r := RSSEnclosure{}
	httpData, httpError := http.Get(os.Getenv("RSS_FEED_URL"))
	if httpError != nil {
		panic(httpError.Error())
	}
	defer httpData.Body.Close()
	xmlContent, readError := ioutil.ReadAll(httpData.Body)
	if readError != nil {
		panic(readError.Error())
	}
	xmlError := xml.Unmarshal(xmlContent, &r)
	if xmlError != nil {
		panic(xmlError.Error())
	}
	fmt.Println(r.LastBuildDate)
	fmt.Println(os.Getenv("LAST_BUILD_DATE"))
	if r.LastBuildDate != os.Getenv("LAST_BUILD_DATE") {
		os.Setenv("LAST_BUILD_DATE", r.LastBuildDate)
		fmt.Println("RSS feed has changed. Willtrigger.")
	}
	// for _, item := range r.ItemList {
	//   //Loop RSS feeds for <item>, which is a post
	//   fmt.Println(item.Title)
	// }
}
