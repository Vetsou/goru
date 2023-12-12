# Goru
Simple web scraper written in Go to extract tags from booru sites.

Supports the following sites:
- Safebooru
- Danbooru
- Gelbooru

## Usage
```bash
go run main.go -site=<SITE> -ids=<ID> -type=<TAGS_TYPE>
```

## Example:
```bash
# Download tags from image with id=4 from safebooru
go run main.go -site=safebooru -ids=4
# Download tags from images with id=4, 22 and 6 from safebooru
go run main.go -site=danbooru -ids="4,22,6"
# Download tags from images with id from 1 to 7 from safebooru
go run main.go -site=safe -ids="1-7"
# Download General and Artist tags from images with id from 1 to 4 and 9 from danbooru
go run main.go -site=dan -ids="1-4,9" -type="ar,g"
```