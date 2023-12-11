package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Virees/goru/flags"
	"github.com/gocolly/colly/v2"
)

const (
	SAFEBOORU_TAGS_LOCATION = "#tag-sidebar > .tag > a"
	DANBOORU_TAGS_LOCATION  = "section#tag-list .search-tag"
	GELBOORU_TAGS_LOCATION  = "ul#tag-list > li[class^='tag-type'] > a"
)

func onResponse(res *colly.Response) {
	fmt.Printf("Got a response from: %s\n", res.Request.URL)
}

func onError(res *colly.Response, e error) {
	fmt.Printf("TagsCollector: %s entering site %s\n (%d)", e, res.Request.URL, res.StatusCode)
}

func setupTagsCollector() *colly.Collector {
	tagsCollector := colly.NewCollector(
		colly.AllowedDomains("safebooru.org", "danbooru.donmai.us", "gelbooru.com"),
	)

	// Config
	tagsCollector.SetRequestTimeout(30 * time.Second)
	tagsCollector.OnResponse(onResponse)
	tagsCollector.OnError(onError)

	onHtmlResponse := func(htmlElem *colly.HTMLElement) {
		fmt.Printf("%s, ", htmlElem.Text)
	}

	tagsCollector.OnHTML(SAFEBOORU_TAGS_LOCATION, onHtmlResponse)
	tagsCollector.OnHTML(DANBOORU_TAGS_LOCATION, onHtmlResponse)
	tagsCollector.OnHTML(GELBOORU_TAGS_LOCATION, onHtmlResponse)

	return tagsCollector
}

func main() {
	flags, err := flags.ParseInputFlags()
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}

	tagsColly := setupTagsCollector()

	urls := flags.GetUrls()
	for _, url := range urls {
		tagsColly.Visit(url)
		if err != nil {
			fmt.Printf("Error visiting the site: %v\n", err)
		}
	}
}
