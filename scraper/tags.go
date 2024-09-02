package scraper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Vetsou/goru/flags"
	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/google/uuid"
)

func SetupTagsCollector(flags flags.GoruFlags) *colly.Collector {
	tagsColly := colly.NewCollector(
		colly.AllowedDomains("safebooru.org", "danbooru.donmai.us", "gelbooru.com"),
		colly.Async(true),
	)

	// Set config
	tagsColly.SetRequestTimeout(30 * time.Second)
	tagsColly.Limit(&colly.LimitRule{Parallelism: 4, DomainGlob: "*"})

	// Handlers
	tagsColly.OnResponse(onResponse)
	tagsColly.OnError(onError)
	tagsColly.OnScraped(onScraped)

	switch flags.Site.Name {
	case "safe", "safebooru":
		tagsColly.OnHTML(SAFE_TAGS_CONTAINER, setupOnTags(SAFE_TAGS_LOCATION, flags.TagsTypeList))
	case "dan", "danbooru":
		tagsColly.OnHTML(DAN_TAGS_CONTAINER, setupOnTags(DAN_TAGS_LOCATION, flags.TagsTypeList))
	case "gel", "gelbooru":
		tagsColly.OnHTML(GEL_TAGS_CONTAINER, setupOnTags(GEL_TAGS_LOCATION, flags.TagsTypeList))
	}

	return tagsColly
}

// Html handler
func setupOnTags(tagsLocation map[string]string, reqTagTypes flags.TagsType) func(*colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		if len(e.Request.Headers.Values("Referer")) != 0 {
			return
		}

		// Get output folder path
		outDirPath := string(e.Request.Ctx.Get("outFolder"))

		// Extract tags
		extractedTags, err := parseTags(tagsLocation, reqTagTypes, e)
		if err != nil {
			fmt.Printf(color.YellowString("Parse tags error: %s\n"), err)
			return
		}

		// Open/Create file
		file, err := createTagsFile(outDirPath, uuid.New().String())
		if err != nil {
			fmt.Printf(color.YellowString("File create error: %s"), err)
			return
		}
		defer file.Close()

		// Save tags to file
		file.WriteString(strings.Join(extractedTags, ", "))
	}
}

// Response handles
func onResponse(res *colly.Response) {
	referer := res.Request.Headers.Values("Referer")
	if len(referer) != 0 {
		fmt.Printf(color.RedString("TagsCollector: Response URL is redirected or not found. Tags will not be downloaded for URL: %s\n"), referer[0])
		return
	}

	fmt.Printf(color.GreenString("TagsCollector: Got a response from: %s (HTTP Code: %d)\n"), res.Request.URL, res.StatusCode)
}

func onError(res *colly.Response, e error) {
	fmt.Printf(color.RedString("TagsCollector: %s entering site %s (HTTP Code: %d)\n"), e, res.Request.URL, res.StatusCode)
}

func onScraped(res *colly.Response) {

}

func createTagsFile(path string, name string) (*os.File, error) {
	outFilePath := filepath.Join(path, name+".txt")

	file, err := os.Create(outFilePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func parseTags(tagsLoc map[string]string, reqTagTypes flags.TagsType, container *colly.HTMLElement) ([]string, error) {
	var extractedTags []string

	for _, tagToDownload := range reqTagTypes {
		foundTags := container.ChildTexts(tagsLoc[tagToDownload])
		extractedTags = append(extractedTags, foundTags...)
	}

	if len(extractedTags) == 0 {
		return nil, errors.New("no tags found")
	}

	return extractedTags, nil
}
