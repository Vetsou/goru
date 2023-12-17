package main

import (
	"fmt"
	"os"

	"github.com/Virees/goru/flags"
	"github.com/Virees/goru/scraper"
	"github.com/gocolly/colly/v2"
)

func main() {
	inputFlags, err := flags.LoadInputFlags()
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}

	tagsColly := scraper.SetupTagsCollector(*inputFlags)

	// Colly request context
	ctx := colly.NewContext()
	ctx.Put("outFolder", string(inputFlags.OutputFolder))

	urls := inputFlags.GetUrls()
	for _, url := range urls {
		tagsColly.Request("GET", url, nil, ctx, nil)
		if err != nil {
			fmt.Printf("Error visiting the site: %v\n", err)
		}
	}

	tagsColly.Wait()
}
