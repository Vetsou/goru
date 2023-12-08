package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

const (
	SAFEBOORU_TAGS_LOCATION = "#tag-sidebar > .tag > a"
	DANBOORU_TAGS_LOCATION  = "section#tag-list .search-tag"
	GELBOORU_TAGS_LOCATION  = "ul#tag-list > li[class^='tag-type'] > a"
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
	handleTags := func(htmlTags *colly.HTMLElement) {
		tags := strings.Fields(htmlTags.Text)
		fmt.Printf("%s\n", tags)
	}

	c.OnHTML(SAFEBOORU_TAGS_LOCATION, handleTags)
	c.OnHTML(DANBOORU_TAGS_LOCATION, handleTags)
	c.OnHTML(GELBOORU_TAGS_LOCATION, handleTags)
}
