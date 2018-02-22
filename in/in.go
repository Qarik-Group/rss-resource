package main

import (
	"encoding/json"
	"encoding/xml"
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

func main() {
	if len(os.Args) < 2 {
		panic("please supply more arguments")
	}
	r := RSSEnclosure{}
	r = r
	httpData, _ := http.Get("https://www.starkandwayne.com/blog/rss/")
	var returnposts []Item
	desireddirectory := "./" + os.Args[1] + "/"
	defer httpData.Body.Close()
	xmlContent, _ := ioutil.ReadAll(httpData.Body)
	xml.Unmarshal(xmlContent, &r)
	returnposts = r.ItemList
	postjson, _ := json.Marshal(returnposts)
	for _, item := range r.ItemList {
		post, _ := json.Marshal(item)
		err := ioutil.WriteFile(desireddirectory+strings.Replace(item.Title, " ", "", -1), post, 0644)
		if err != nil {
			panic(err.Error())
		}
	}
	postjson = postjson
	//fmt.Println(postjson)
	//os.Stdout.Write(postjson)
}
