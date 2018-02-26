package rss

import (
	"encoding/xml"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`

	Title         string `xml:"channel>title"`
	Link          string `xml:"channel>link"`
	Description   string `xml:"channel>description"`
	LastBuildDate string `xml:"channel>lastBuildDate"`
	PubDate       string `xml:"channel>pubDate"`
	Items         []struct {
		Title       string   `xml:"title"`
		Author      string   `xml:"author"`
		Link        string   `xml:"link"`
		Description string   `xml:"description"`
		Content     string   `xml:"encoded"`
		PubDate     string   `xml:"pubDate"`
		Comments    string   `xml:"comments"`
		Category    []string `xml:"category"`
	} `xml:"channel>item"`
}
