package flags

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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

// Flag representing the path to the folder where image tags will be saved
type OutputPath string

func (p *OutputPath) String() string {
	return fmt.Sprintf("%v", *p)
}

func (p *OutputPath) Set(value string) error {
	absPath, err := filepath.Abs(value)
	if err != nil {
		return fmt.Errorf("path invalid: %s", err)
	}

	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", absPath)
	} else if err != nil {
		return fmt.Errorf("error checking path: %s", err)
	}

	*p = OutputPath(absPath)
	return nil
}
