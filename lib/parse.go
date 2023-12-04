package lib

import (
	"errors"
	"strings"
)

func ParseBooruTags(tagsStr string) ([]string, error) {
	tags := strings.Fields(tagsStr)

	if len(tags) > 0 {
		tags = tags[:len(tags)-1]
		return tags, nil
	}
	return nil, errors.New("no tags found")
}
