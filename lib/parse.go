package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// Wrapper for the Colly collector, designed for extracting tags
type TagsCollector struct {
	handle *colly.Collector
}

func NewTagsCollector() *TagsCollector {
	c := colly.NewCollector(
		colly.AllowedDomains("safebooru.org", "danbooru.donmai.us", "gelbooru.com"),
	)
	c.SetRequestTimeout(30 * time.Second)

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from: ", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("TagsCollector: %s entering site %s: ", e, r.Request.URL)
	})

	setupOnHTML(c)
	return &TagsCollector{handle: c}
}

func (tc *TagsCollector) Visit(url string) error {
	return tc.handle.Visit(url)
}

func setupOnHTML(c *colly.Collector) {
	// Safebooru
	c.OnHTML("ul#tag-sidebar", func(e *colly.HTMLElement) {
		fmt.Printf("%s\n", parseTags(e.Text))
	})

	// Danbooru
	c.OnHTML("section#tag-list", func(e *colly.HTMLElement) {
		fmt.Printf("%s\n", parseTags(e.Text))
	})

	// Gelbooru
	c.OnHTML("ul#tag-list", func(e *colly.HTMLElement) {
		fmt.Printf("%s\n", parseTags(e.Text))
	})
}

// Parse tags
func parseTags(tagsStr string) []string {
	tags := strings.Fields(tagsStr)

	if len(tags) > 0 {
		tags = tags[:len(tags)-1]
	}
	return tags
}
