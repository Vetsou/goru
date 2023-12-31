# Goru
Simple web scraper written in Go to extract tags from booru sites.

Supports the following sites:
- Safebooru
- Danbooru
- Gelbooru

## Usage
```bash
go run main.go -site=<SITE> -ids=<ID> -type=<TAGS_TYPE>
  - SITE - Source image site ["danbooru"|"dan"|"safebooru"|"safe"|"gelbooru"|"gel"]
  - ID - Comma-separated list of image IDs ["2,4,5"|"1-22"|"13-7"|"6,8-12,22,44-55"]
  - TAGS_TYPE - Comma-separated list of tag types to include ["a"|"cr,ch"|"ar,g,md"]
    - a - All tags
    - cr - Copyright tags
    - ch - Character tags
    - ar - Artist tags
    - g - General tags
    - md - Metadata tags
```

## Example:
```bash
# Download tags from image with id=4 from safebooru. Save result into "./output" folder.
go run main.go -site=safebooru -ids=4 -out="output"
# Download tags from images with id=4, 22 and 6 from safebooru. Save result into "./output/subfolder" folder.
go run main.go -site=danbooru -ids="4,22,6" -out="output/subfolder"
# Download tags from images with id from 1 to 7 from safebooru. Save result into "C:\img" folder.
go run main.go -site=safe -ids="1-7" -out="C:\imgs"
# Download General and Artist tags from images with id from 1 to 4 and 9 from danbooru. Save result into "C:\img\sub" folder.
go run main.go -site=dan -ids="1-4,9" -type="ar,g" -out="C:\img\sub"
```