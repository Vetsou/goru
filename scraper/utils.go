package scraper

import (
	"errors"
	"os"
	"path/filepath"

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

func CreateFile(folderPath string, id string) (*os.File, error) {
	outFilePath := filepath.Join(folderPath, id+".txt")

	file, err := os.Create(outFilePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
