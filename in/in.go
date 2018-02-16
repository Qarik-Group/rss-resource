package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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
	Category    []string      `xml:"category"`
}

func main() {
	r := RSSEnclosure{}
	r = r
	httpData, _ := http.Get("https://www.starkandwayne.com/blog/rss/")
	var returnposts []Item
	// if httpError != nil {
	// 	panic(httpError.Error())
	// }
	defer httpData.Body.Close()
	xmlContent, _ := ioutil.ReadAll(httpData.Body)
	// if readError != nil {
	// 	panic(readError.Error())
	// }
	xml.Unmarshal(xmlContent, &r)
	//fmt.Println(r)
	// xmlError := xml.Unmarshal(xmlContent, &r)
	// if xmlError != nil {
	// 	panic(xmlError.Error())
	// }
	// fmt.Println(r.LastBuildDate)
	// fmt.Println(os.Getenv("LAST_BUILD_DATE"))
	// if r.LastBuildDate != os.Getenv("LAST_BUILD_DATE") {
	// 	os.Setenv("LAST_BUILD_DATE", r.LastBuildDate)
	// 	fmt.Println("RSS feed has changed. Willtrigger.")
	// }
	for _, item := range r.ItemList {
		//Loop RSS feeds for <item>, which is a post
		curItem := item
		for _, categ := range curItem.Category {
			if categ == "SHIELD" {
				returnposts = append(returnposts, item)
			}
		}
	}
	postjson, _ := json.Marshal(returnposts)
	for _, post := range returnposts {
		fmt.Println(post.Title)
	}
	fmt.Println(postjson)

}
