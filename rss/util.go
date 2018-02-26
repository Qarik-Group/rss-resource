package rss

import (
	"fmt"
	"regexp"
	"strings"
)

func filename(date, title string) string {
	return regexp.MustCompile(`^-|-$`).ReplaceAllString(
		regexp.MustCompile(`--+`).ReplaceAllString(
			regexp.MustCompile(`[^a-z0-9]`).ReplaceAllString(
				strings.ToLower(
					fmt.Sprintf("%s-%s", date, title),
				),
				"-",
			),
			"",
		),
		"",
	)
}
