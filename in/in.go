package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

type ReturnJson struct {
	Version  Version `json:"version"`
	Metadata []Meta  `json:"metadata"`
}

type Version struct {
	Ref string `json:"ref"`
}

type Meta struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	if len(os.Args) < 2 {
		panic("please supply more arguments")
	}
	r := RSSEnclosure{}
	r = r
	httpData, _ := http.Get("https://www.starkandwayne.com/blog/rss/")
	desireddirectory := os.Args[1]
	os.MkdirAll(desireddirectory, 0777) //0644? note should probably change permissions
	defer httpData.Body.Close()
	xmlContent, _ := ioutil.ReadAll(httpData.Body)
	xml.Unmarshal(xmlContent, &r)
	for _, item := range r.ItemList {
		post, _ := json.Marshal(item)
		err := ioutil.WriteFile(desireddirectory+"/"+strings.Replace(item.Title, " ", "", -1)+".json", post, 0777)
		if err != nil {
			panic(err.Error())
		}
	}
	rawVersion := Version{Ref: "latest"}
	rawMeta := Meta{Name: "Latest Title", Value: r.ItemList[0].Title}
	metalist := []Meta{rawMeta}
	rawReturn := ReturnJson{Version: rawVersion, Metadata: metalist}
	json, _ := json.Marshal(rawReturn)
	fmt.Printf("%s\n", json)
}
