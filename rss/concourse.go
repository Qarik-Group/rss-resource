package rss

type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Output struct {
	Version struct {
		Ref string `json:"ref"`
	} `json:"version"`

	Metadata []Metadata `json:"metadata"`
}
