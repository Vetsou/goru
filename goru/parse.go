package goru

import (
	"errors"
	"flag"
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

func ParseInputFlags() int {
	id := flag.Int("id", 0, "Image id")
	flag.Parse()

	return *id
}
