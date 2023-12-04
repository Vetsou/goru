# Goru
Simple web scrapper written in go to extract tags from booru sites.

Supports the following sites:
- Safebooru
- Danbooru
- Gelbooru

## Usage
```ps
go run main.go -site <SITE_NAME (safebooru|gelbooru|danbooru)> -id <PICTURE_ID>
```

## Example:
```ps
go run main.go -site safebooru -id 4634700
```