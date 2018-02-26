package rss

type Config struct {
	Version struct {
		Ref string `json:"ref"`
	} `json:"version"`

	Source struct {
		URL           string `json:"url"`
		SkipTLSVerify bool   `json:"skip_tls_verify"`
	} `json:"source"`
}
