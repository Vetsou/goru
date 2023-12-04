# Goru
Simple web scrapper written in go to extract tags from booru sites.

Supports the following sites:
- Safebooru
- Danbooru
- Gelbooru

## Usage
```
go run main.go -site <SITE_NAME (safebooru|gelbooru|danbooru)> -id <PICTURE_ID>
```

## Example:
```
go run main.go -site safebooru -id 4634700
```