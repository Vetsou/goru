package flags

import (
	"errors"
	"fmt"
	"strings"
)

// Map of tag types
var allowedTagTypes = map[string]bool{
	"a":  true, // All
	"cr": true, // Copyright
	"ch": true, // Character
	"ar": true, // Artist
	"g":  true, // General
	"md": true, // Metadata
}

// Flag representing the list of tags that should be downloaded
type TagsType []string

func (t *TagsType) String() string {
	return fmt.Sprintf("%v", *t)
}

func (t *TagsType) Set(value string) error {
	tagTypes := strings.Split(value, ",")

	for _, tagType := range tagTypes {
		if !allowedTagTypes[tagType] {
			return errors.New("unsupported tag type")
		}

		// If 'a' (All) tag detected return all tags list
		if tagType == "a" {
			*t = TagsType{"cr", "ch", "ar", "g", "md"}
			return nil
		}
	}

	*t = tagTypes
	return nil
}
