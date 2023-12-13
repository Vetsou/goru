package scraper

import (
	"fmt"
	"strings"
	"time"

	"github.com/Virees/goru/flags"
	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
)

func SetupTagsCollector(flags flags.GoruFlags) *colly.Collector {
	tagsColly := colly.NewCollector(
		colly.AllowedDomains("safebooru.org", "danbooru.donmai.us", "gelbooru.com"),
	)

	// Set config
	tagsColly.SetRequestTimeout(30 * time.Second)
	tagsColly.OnResponse(onResponse)
	tagsColly.OnError(onError)
	tagsColly.OnScraped(onScraped)

	switch flags.Site.Name {
	case "safe", "safebooru":
		tagsColly.OnHTML(SAFE_TAGS_CONTAINER, setupOnSafebooruTags(flags.TagsTypeList))
	case "dan", "danbooru":
		tagsColly.OnHTML(DAN_TAGS_CONTAINER, setupOnDanbooruTags(flags.TagsTypeList))
	case "gel", "gelbooru":
		tagsColly.OnHTML(GEL_TAGS_CONTAINER, setupOnGelbooruTags(flags.TagsTypeList))
	}

	return tagsColly
}

// Html handles
func setupOnDanbooruTags(tagsToDownload flags.TagsType) func(*colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		var extractedTags []string

		for _, tagToDownload := range tagsToDownload {
			foundTags := e.ChildTexts(DAN_TAGS_LOCATION[tagToDownload])
			extractedTags = append(extractedTags, foundTags...)
		}

		if len(extractedTags) == 0 {
			fmt.Println("No tags found")
			return
		}

		fmt.Printf(color.GreenString("%s"), strings.Join(extractedTags, ", "))
	}
}

func setupOnSafebooruTags(tagsToDownload flags.TagsType) func(*colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		var extractedTags []string

		for _, tagToDownload := range tagsToDownload {
			foundTags := e.ChildTexts(SAFE_TAGS_LOCATION[tagToDownload])
			extractedTags = append(extractedTags, foundTags...)
		}

		if len(extractedTags) == 0 {
			fmt.Println("No tags found")
			return
		}

		fmt.Printf(color.GreenString("%s"), strings.Join(extractedTags, ", "))
	}
}

func setupOnGelbooruTags(tagsToDownload flags.TagsType) func(*colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		var extractedTags []string

		for _, tagToDownload := range tagsToDownload {
			foundTags := e.ChildTexts(GEL_TAGS_LOCATION[tagToDownload])
			extractedTags = append(extractedTags, foundTags...)
		}

		if len(extractedTags) == 0 {
			fmt.Println("No tags found")
			return
		}

		fmt.Printf(color.GreenString("%s"), strings.Join(extractedTags, ", "))
	}
}

// Response handles
func onResponse(res *colly.Response) {
	fmt.Printf(color.YellowString("\n\nGot a response from: %s\n"), res.Request.URL)
}

func onError(res *colly.Response, e error) {
	fmt.Printf(color.RedString("\n\nTagsCollector: %s entering site %s (Code: %d)"), e, res.Request.URL, res.StatusCode)
}

func onScraped(res *colly.Response) {

}
