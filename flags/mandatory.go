package flags

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Empty flag error
var ErrEmptyFlag = errors.New("mandatory flag is empty")

// Map structure that contains URL templates for supported sites
var SUPPORTED_SITES = map[string]string{
	"safe":      "https://safebooru.org/index.php?page=post&s=view&id=%d",
	"safebooru": "https://safebooru.org/index.php?page=post&s=view&id=%d",
	"dan":       "https://danbooru.donmai.us/posts/%d",
	"danbooru":  "https://danbooru.donmai.us/posts/%d",
	"gel":       "https://gelbooru.com/index.php?page=post&s=view&id=%d",
	"gelbooru":  "https://gelbooru.com/index.php?page=post&s=view&id=%d",
}

// Function parsing id range (e.g. 1-2, 3-21, 70-23)
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

	if fromId > toId {
		fromId, toId = toId, fromId
	}

	for i := fromId; i <= toId; i++ {
		*idList = append(*idList, i)
	}

	return nil
}

// Flag representing the list of image IDs
type IDList []int

func (idList *IDList) String() string {
	return fmt.Sprintf("%v", *idList)
}

func (idList *IDList) Set(value string) error {
	if value == "" {
		return ErrEmptyFlag
	}

	idStrings := strings.Split(value, ",")

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
			return fmt.Errorf("error parsing id: %v", err)
		}
		*idList = append(*idList, id)
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
		return ErrEmptyFlag
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
