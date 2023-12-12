package main

import (
	"fmt"
	"os"

	"github.com/Virees/goru/flags"
	"github.com/Virees/goru/scraper"
)

func main() {
	inputFlags, err := flags.LoadInputFlags()
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}

	tagsColly := scraper.SetupTagsCollector(*inputFlags)

	urls := inputFlags.GetUrls()
	for _, url := range urls {
		tagsColly.Visit(url)
		if err != nil {
			fmt.Printf("Error visiting the site: %v\n", err)
		}
	}
}
