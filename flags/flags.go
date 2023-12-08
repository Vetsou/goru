package flags

import (
	"flag"
	"fmt"
)

type GoruFlags struct {
	Site   SourceSite
	IdList IDList
}

func ParseInputFlags() (*GoruFlags, error) {
	var site SourceSite
	var idList IDList

	siteFlag := flag.String("site", "", "Source image site")
	idsFlag := flag.String("ids", "", "Comma-separated list of IDs")
	flag.Parse()

	if err := site.Set(*siteFlag); err != nil {
		return nil, fmt.Errorf("set site flag failed: %w", err)
	}

	if err := idList.Set(*idsFlag); err != nil {
		return nil, fmt.Errorf("set ids flag failed: %w", err)
	}

	gf := &GoruFlags{
		Site:   site,
		IdList: idList,
	}

	return gf, nil
}

func (gflags *GoruFlags) GetUrls() []string {
	var urls []string

	for _, id := range gflags.IdList {
		urls = append(urls, fmt.Sprintf(gflags.Site.UrlTemplate, id))
	}

	return urls
}
