package main

import (
	"fmt"
	"time"

	"github.com/Virees/goru/lib"
	"github.com/gocolly/colly/v2"
)

func setupColly() *colly.Collector {
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

	// Parse Danbooru tags
	c.OnHTML("li[data-tag-name]", func(e *colly.HTMLElement) {
		fmt.Printf("%s\n", e.Attr("data-tag-name"))
	})

	// Parse Safebooru/Gelbooru tags
	c.OnHTML("a[href^='index.php?page=post&s=list&tags=']", func(e *colly.HTMLElement) {
		fmt.Printf("%s\n", e.Text)
	})

	return c
}

func main() {
	flags := lib.ParseInputFlags()
	c := setupColly()

	url, err := flags.GetUrl()
	if err != nil {
		fmt.Println("Error generating URL: ", err)
	}

	err = c.Visit(url)
	if err != nil {
		fmt.Println("Error visiting the site: ", err)
	}
}
