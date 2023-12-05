package lib

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

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
		fmt.Printf("Error %s entering site %s: ", e, r.Request.URL)
	})

	tc := &TagsCollector{handle: c}
	tc.setupTagParsers()

	return tc
}

func (tc *TagsCollector) Visit(url string) error {
	return tc.handle.Visit(url)
}

func parseBooruTags(tagsStr string) ([]string, error) {
	tags := strings.Fields(tagsStr)

	if len(tags) > 0 {
		tags = tags[:len(tags)-1]
		return tags, nil
	}
	return nil, errors.New("no tags found")
}

func (tc *TagsCollector) setupTagParsers() {
	// Parse Danbooru tags
	tc.handle.OnHTML("li[data-tag-name]", func(e *colly.HTMLElement) {
		fmt.Printf("%s\n", e.Attr("data-tag-name"))
	})

	// Parse Safebooru/Gelbooru tags
	tc.handle.OnHTML("a[href^='index.php?page=post&s=list&tags=']", func(e *colly.HTMLElement) {
		tags, err := parseBooruTags(e.Text)
		if err != nil {
			fmt.Println("Error parsing tags: ", e)
			return
		}

		fmt.Printf("%s\n", tags)
	})
}
