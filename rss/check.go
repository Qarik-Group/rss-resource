package rss

import (
	"crypto/sha1"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func Check() {
	var cfg Config
	out := make([]struct {
		Ref string `json:"ref"`
	}, 1)

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read configuration from standard input: %s\n", err)
		os.Exit(1)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to unmarshal configuration: %s\n", err)
		os.Exit(1)
	}

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: cfg.Source.SkipTLSVerify,
			},
		},
	}

	if cfg.Source.URL == "" {
		fmt.Fprintf(os.Stderr, "The (required) source.url parameter was not specified.\n")
		os.Exit(1)
	}

	res, err := c.Get(cfg.Source.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to retrieve %s: %s\n", cfg.Source.URL, err)
		os.Exit(1)
	}

	rss, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read response from %s: %s\n", cfg.Source.URL, err)
		os.Exit(1)
	}

	re := regexp.MustCompile("<lastBuildDate>.*</lastBuildDate>")
	rssString := strings.Replace(string(rss), re.FindString(string(rss)), "", -1)

	out[0].Ref = fmt.Sprintf("%x", sha1.Sum([]byte(rssString)))
	b, err = json.Marshal(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal output JSON: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(b))
}
