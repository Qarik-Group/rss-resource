package rss

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func In() {
	var (
		r   RSS
		out Output
		cfg Config
	)

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
	err = ioutil.WriteFile("feed.xml", rss, 0666)
	if err != nil {
		fmt.Printf("Failed to save feed.xml to local directory: %s\n", err)
		os.Exit(1)
	}

	err = xml.Unmarshal(rss, &r)
	if err != nil {
		fmt.Printf("Failed to parse Atom/RSS feed xml: %s\n", err)
		os.Exit(1)
	}

	os.Mkdir("posts", 0777)
	for i, item := range r.Items {
		b, err = json.Marshal(item)
		if err != nil {
			fmt.Printf("Failed to marshal post #%d (%s) to JSON: %s\n", i+1, item.Title, err)
			os.Exit(1)
		}
		file := fmt.Sprintf("posts/%s.json", filename(item.PubDate, item.Title))
		err = ioutil.WriteFile(file, b, 0666)
		if err != nil {
			fmt.Printf("Failed to write post #%d (%s) JSON to %s: %s\n", i+1, item.Title, file, err)
			os.Exit(1)
		}
	}

	out.Version.Ref = "(none)"
	if len(r.Items) > 0 {
		out.Version.Ref = r.Items[0].PubDate
		out.Metadata = append(out.Metadata, Metadata{Name: "Title:", Value: r.Items[0].Title})
		out.Metadata = append(out.Metadata, Metadata{Name: "Author:", Value: r.Items[0].Author})
		out.Metadata = append(out.Metadata, Metadata{Name: "Link:", Value: r.Items[0].Link})
	}

	b, err = json.Marshal(out)
	if err != nil {
		fmt.Printf("Failed to marshal output JSON: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(b))
}
