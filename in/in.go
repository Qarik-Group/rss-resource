package main

import (
  "io/ioutil"
  "encoding/xml"
  "html/template"
  "net/http"
  "fmt"
)

type RSSEnclosure struct {
  XMLName xml.Name `xml:"rss"`
  Version string `xml:"version,attr"`

  Title string `xml:"channel>title"`
  Link string `xml:"channel>link"`
  Description string  `xml:"channel>description"`

  PubDate string `xml:"channel>pubDate"`
  ItemList []Item `xml:"channel>item"`
}

type Item struct {

  Title string `xml:"title"`
  Link string `xml:"link"`
  Description template.HTML `xml:"description"`

  Content template.HTML`xml:"encoded"`
  PubDate string `xml:"pubDate"`
  Comments string `xml:"comments"`
}


func main() {
  r := RSSEnclosure{}
  httpData, httpError := http.Get("https://starkandwayne.com/blog/rss/")
  if httpError != nil { panic(httpError.Error()) }
  defer httpData.Body.Close()
  xmlContent, readError := ioutil.ReadAll(httpData.Body)
  if readError != nil { panic(readError.Error()) }
  xmlError := xml.Unmarshal(xmlContent, &r)
  if xmlError != nil { panic(xmlError.Error()) }
  for _, item := range r.ItemList {
  fmt.Println(item.Title)
  }
}
