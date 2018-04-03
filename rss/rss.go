package rss

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

type Atom struct {
	XMLName xml.Name `xml:"feed"`

	Title   string `xml:"title"`
	Updated string `xml:"updated"`

	Entries []struct {
		ID      string `xml:"id"`
		Updated string `xml:"updated"`
		Title   string `xml:"title"`
		Content string `xml:"content"`
		Author  string `xml:"author>name"`
	} `xml:"entry"`
}

type Post struct {
	Title       string `json:"Title"`
	Intro       string `json:"Intro"`
	Full        string `json:"Content"`
	Timestamp   int64  `json:"PubDate"`
	Author      string `json:"Author"`
	Description string `json:"Description"`
	Link        string `json:"Link"`
}

type Feed []Post

func ParseURL(u string, verify bool) (Feed, []byte, error) {
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !verify,
			},
		},
	}

	res, err := c.Get(u)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to retrieve %s: %s", u, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to read response from %s: %s", u, err)
	}

	f, err := Parse(body)
	if err != nil {
		return nil, nil, err
	}

	return f, body, nil
}

func Parse(b []byte) (Feed, error) {
	if f, err := ParseRSS(b); err == nil {
		return f, nil
	}
	if f, err := ParseAtom(b); err == nil {
		return f, nil
	}
	return nil, fmt.Errorf("Failed to parse feed XML: doesn't look like either RSS or Atom")
}

func ParseRSS(b []byte) (Feed, error) {
	var r RSS
	if err := xml.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	f := make(Feed, 0)
	for _, item := range r.Items {
		ts, err := time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", item.PubDate)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse RSS pubDate '%s': %s", item.PubDate, err)
		}
		f = append(f, Post{
			Title:       item.Title,
			Intro:       item.Description,
			Full:        item.Content,
			Timestamp:   ts.Unix(),
			Author:      item.Author, // FIXME
			Description: item.Description,
			Link:        item.Link,
		})
	}
	return f, nil
}

func ParseAtom(b []byte) (Feed, error) {
	var a Atom
	if err := xml.Unmarshal(b, &a); err != nil {
		return nil, err
	}

	f := make(Feed, 0)
	for _, item := range a.Entries {
		ts, err := time.Parse("2006-01-02T15:04:05-07:00", item.Updated)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse Atom timestamp '%s': %s", item.Updated, err)
		}
		f = append(f, Post{
			Title:     item.Title,
			Intro:     item.Content, // FIXME
			Full:      item.Content,
			Timestamp: ts.Unix(),
			Author:    item.Author,
		})
	}
	return f, nil
}
