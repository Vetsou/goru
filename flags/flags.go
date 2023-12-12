package flags

import (
	"flag"
	"fmt"
)

type GoruFlags struct {
	Site         SourceSite
	IdList       IDList
	TagsTypeList TagsType
}

func LoadInputFlags() (*GoruFlags, error) {
	var site SourceSite
	var idList IDList
	var tagsTypeList TagsType

	// Read mandatory input flags
	siteFlag := flag.String("site", "", "Source image site (e.g., safebooru, safe)")
	idsFlag := flag.String("ids", "", "Comma-separated list of image IDs")
	tagTypes := flag.String("type", "", "Comma-separated list of tag types to include (e.g., all, copyright, character)")
	flag.Parse()

	// Parse mandatory flags
	if err := site.Set(*siteFlag); err != nil {
		return nil, fmt.Errorf("set site flag failed: %w", err)
	}

	if err := idList.Set(*idsFlag); err != nil {
		return nil, fmt.Errorf("set ids flag failed: %w", err)
	}

	// Parse optional flags
	if *tagTypes != "" {
		if err := tagsTypeList.Set(*tagTypes); err != nil {
			return nil, fmt.Errorf("set tags type flag failed: %w", err)
		}
	} else {
		tagsTypeList.Set("a")
	}

	return &GoruFlags{
		Site:         site,
		IdList:       idList,
		TagsTypeList: tagsTypeList,
	}, nil
}

func (gflags *GoruFlags) GetUrls() []string {
	var urls []string

	for _, id := range gflags.IdList {
		urls = append(urls, fmt.Sprintf(gflags.Site.UrlTemplate, id))
	}

	return urls
}
