package main

import ("encoding/xml")
import ("fmt")
import ("os")
import ("io/ioutil")


type RSSEnclousure struct {

}

type Blog struct {
  Post Item
  PostList []Item `xml:"item>"`
}

type Item struct {
  Title string `xml:"title>"`
  Link string
  Category map[string] bool
  Description string
  Content string
}



func (post Item) String() string {
  return fmt.Sprintf("%s - %d", post.Title, post.Category)
}

func main() {
  xmlFile, err := os.Open("rss.xml")
  if err != nil {
    fmt.Println("Error opening file: ", err)
    return
  }
  defer xmlFile.Close()

  var blog Blog
  b, _ := ioutil.ReadAll(xmlFile)
    xml.Unmarshal(b, &blog)

  fmt.Println(blog.Post)
  for _, post := range blog.PostList {
    fmt.Printf("==>", post)
  }
}
