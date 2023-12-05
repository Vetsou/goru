package lib

import (
	"errors"
	"flag"
	"fmt"
)

type Flags struct {
	Site string
	Id   int
}

func ParseInputFlags() (Flags, error) {
	site := flag.String("site", "", "Source image site")
	id := flag.Int("id", 0, "Source image id")
	flag.Parse()

	flags := Flags{
		Site: *site,
		Id:   *id,
	}

	return flags, nil
}

func (f *Flags) GetUrl() (string, error) {
	var supportedSites = map[string]string{
		"safebooru": "https://safebooru.org/index.php?page=post&s=view&id=%d",
		"danbooru":  "https://danbooru.donmai.us/posts/%d",
		"gelbooru":  "https://gelbooru.com/index.php?page=post&s=view&id=%d",
	}

	urlTemplate, ok := supportedSites[f.Site]
	if !ok {
		return "", errors.New("unsupported source site")
	}

	return fmt.Sprintf(urlTemplate, f.Id), nil
}
