package flags

import (
	"flag"
	"fmt"
)

type GoruFlags struct {
	Site         SourceSite
	IdList       IDList
	TagsTypeList TagsType
	OutputFolder OutputPath
}

func LoadInputFlags() (*GoruFlags, error) {
	var site SourceSite
	var idList IDList
	var tagsTypeList TagsType
	var outputLocation OutputPath

	// Read mandatory input flags
	siteFlag := flag.String("site", "", "Source image site (e.g., safebooru, safe)")
	idsFlag := flag.String("ids", "", "Comma-separated list of image IDs")

	// Read optional input flags
	tagTypesFlag := flag.String("type", "a", "Comma-separated list of tag types to include (e.g., all, copyright, character)")
	outFolderFlag := flag.String("out", "./", "Path to the folder where image tags will be saved")
	flag.Parse()

	// Parse mandatory flags
	if err := site.Set(*siteFlag); err != nil {
		return nil, fmt.Errorf("set site flag failed: %w", err)
	}

	if err := idList.Set(*idsFlag); err != nil {
		return nil, fmt.Errorf("set ids flag failed: %w", err)
	}

	// Parse optional flags
	if err := tagsTypeList.Set(*tagTypesFlag); err != nil {
		return nil, fmt.Errorf("set tags type flag failed: %w", err)
	}

	if err := outputLocation.Set(*outFolderFlag); err != nil {
		return nil, fmt.Errorf("set out flag failed: %w", err)
	}

	return &GoruFlags{
		Site:         site,
		IdList:       idList,
		TagsTypeList: tagsTypeList,
		OutputFolder: outputLocation,
	}, nil
}

func (gflags *GoruFlags) GetUrls() []string {
	var urls []string

	for _, id := range gflags.IdList {
		urls = append(urls, fmt.Sprintf(gflags.Site.UrlTemplate, id))
	}

	return urls
}
