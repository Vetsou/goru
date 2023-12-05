package flags

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Map structure that contains URL templates for supported sites
var SUPPORTED_SITES = map[string]string{
	"safebooru": "https://safebooru.org/index.php?page=post&s=view&id=%d",
	"danbooru":  "https://danbooru.donmai.us/posts/%d",
	"gelbooru":  "https://gelbooru.com/index.php?page=post&s=view&id=%d",
}

// Flag representing the list of image IDs
type IDList []int

func (ids *IDList) String() string {
	return fmt.Sprintf("%v", *ids)
}

func (idList *IDList) Set(value string) error {
	idStrings := strings.Split(value, ",")

	if len(idStrings) == 0 {
		return fmt.Errorf("flag ids: is empty")
	}

	for _, idStr := range idStrings {
		// Parse range list (e.g. 1-23)
		if strings.Contains(idStr, "-") {
			err := parseRange(idStr, idList)
			if err != nil {
				return err
			}

			continue
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return fmt.Errorf("flag ids: error parsing id: %v", err)
		}
		*idList = append(*idList, id)
	}

	return nil
}

func parseRange(idStr string, idList *IDList) error {
	rangeParts := strings.Split(idStr, "-")
	if len(rangeParts) != 2 {
		return fmt.Errorf("invalid id range format: %s", idStr)
	}

	fromId, err := strconv.Atoi(rangeParts[0])
	if err != nil {
		return fmt.Errorf("error parsing range value: %v", err)
	}

	toId, err := strconv.Atoi(rangeParts[1])
	if err != nil {
		return fmt.Errorf("error parsing range value: %v", err)
	}

	for i := fromId; i <= toId; i++ {
		*idList = append(*idList, i)
	}

	return nil
}

// Flag representing the image source site
type SourceSite struct {
	Name        string
	UrlTemplate string
}

func (srcSite *SourceSite) String() string {
	return fmt.Sprintf("%s", *srcSite)
}

func (srcSite *SourceSite) Set(value string) error {
	if value == "" {
		return errors.New("sourceSite flag is empty")
	}

	urlTemplate, ok := SUPPORTED_SITES[value]
	if !ok {
		return errors.New("unsupported source site")
	}

	*srcSite = SourceSite{
		Name:        value,
		UrlTemplate: urlTemplate,
	}

	return nil
}
