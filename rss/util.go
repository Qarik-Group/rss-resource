package rss

import (
	"fmt"
	"regexp"
	"strings"
)

func filename(post Post) string {
	return regexp.MustCompile(`^-|-$`).ReplaceAllString(
		regexp.MustCompile(`--+`).ReplaceAllString(
			regexp.MustCompile(`[^a-z0-9]`).ReplaceAllString(
				strings.ToLower(
					fmt.Sprintf("%d-%s", post.Timestamp, post.Title),
				),
				"-",
			),
			"",
		),
		"",
	)
}
