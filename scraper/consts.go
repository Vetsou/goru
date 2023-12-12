package scraper

// Tag container locations
const (
	SAFE_TAGS_CONTAINER = "#tag-sidebar"
	DAN_TAGS_CONTAINER  = "section#tag-list"
	GEL_TAGS_CONTAINER  = "ul#tag-list"
)

// Safebooru input site html elements location
var SAFE_TAGS_LOCATION = map[string]string{
	"cr": ".tag-type-copyright > a",
	"ch": ".tag-type-character > a",
	"ar": ".tag-type-artist > a",
	"g":  ".tag-type-general > a",
	"md": ".tag-type-metadata > a",
}

// Danbooru input site html elements location
var DAN_TAGS_LOCATION = map[string]string{
	"cr": ".tag-type-3 > .search-tag",
	"ch": ".tag-type-4 > .search-tag",
	"ar": ".tag-type-1 > .search-tag",
	"g":  ".tag-type-0 > .search-tag",
	"md": ".tag-type-5 > .search-tag",
}

// Gelbooru input site html elements location
var GEL_TAGS_LOCATION = map[string]string{
	"cr": ".tag-type-copyright > a",
	"ch": ".tag-type-character > a",
	"ar": ".tag-type-artist > a",
	"g":  ".tag-type-general > a",
	"md": ".tag-type-metadata > a",
}
