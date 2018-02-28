package rss

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		fmt.Fprintf(os.Stderr, "Failed to read configuration from standard input: %s\n", err)
		os.Exit(1)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to unmarshal configuration: %s\n", err)
		os.Exit(1)
	}

	if cfg.Source.URL == "" {
		fmt.Fprintf(os.Stderr, "The (required) source.url parameter was not specified.\n")
		os.Exit(1)
	}

	feed, raw, err := ParseURL(cfg.Source.URL, !cfg.Source.SkipTLSVerify)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("feed.xml", raw, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save feed.xml to local directory: %s\n", err)
		os.Exit(1)
	}

	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	root = fmt.Sprintf("%s/posts", root)

	os.Mkdir(root, 0777)
	for i, post := range feed {
		b, err = json.Marshal(post)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to marshal post #%d (%s) to JSON: %s\n", i+1, post.Title, err)
			os.Exit(1)
		}
		file := fmt.Sprintf("%s/%s.json", root, filename(post))
		fmt.Fprintf(os.Stderr, "- writing %s...\n", file)
		err = ioutil.WriteFile(file, b, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write post #%d (%s) JSON to %s: %s\n", i+1, post.Title, file, err)
			os.Exit(1)
		}
	}

	out.Version.Ref = "(none)"
	if len(r.Items) > 0 {
		out.Version.Ref = r.Items[0].PubDate
		out.Metadata = append(out.Metadata, Metadata{Name: "title", Value: r.Items[0].Title})
		out.Metadata = append(out.Metadata, Metadata{Name: "link", Value: r.Items[0].Link})
	}

	b, err = json.Marshal(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal output JSON: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(b))
}
