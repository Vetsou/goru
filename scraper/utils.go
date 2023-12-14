package scraper

import (
	"errors"

	"github.com/Virees/goru/flags"
	"github.com/gocolly/colly/v2"
)

func ParseTags(tagsLocation map[string]string, requestedTags flags.TagsType, rootHtml *colly.HTMLElement) ([]string, error) {
	var extractedTags []string

	for _, tagToDownload := range requestedTags {
		foundTags := rootHtml.ChildTexts(tagsLocation[tagToDownload])
		extractedTags = append(extractedTags, foundTags...)
	}

	if len(extractedTags) == 0 {
		return nil, errors.New("no tags found")
	}

	return extractedTags, nil
}
