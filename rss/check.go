package rss

import (
	"fmt"
	"os"
	"io/ioutil"
	"crypto/sha1"
	"net/http"
	"encoding/json"
	"crypto/tls"
)

func Check() {
	var cfg Config
	out := make([]struct { Ref string `json:"ref"` }, 1)

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Failed to read configuration from standard input: %s\n", err)
		os.Exit(1)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		fmt.Printf("Failed to unmarshal configuration: %s\n", err)
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
		fmt.Printf("The (required) source.url parameter was not specified.\n")
		os.Exit(1)
	}

	res, err := c.Get(cfg.Source.URL)
	if err != nil {
		fmt.Printf("Failed to retrieve %s: %s\n", cfg.Source.URL, err)
		os.Exit(1)
	}

	rss, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Failed to read response from %s: %s\n", cfg.Source.URL, err)
		os.Exit(1)
	}

	out[0].Ref = fmt.Sprintf("%x", sha1.Sum(rss))
	b, err = json.Marshal(out)
	if err != nil {
		fmt.Printf("Failed to marshal output JSON: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(b))
}
