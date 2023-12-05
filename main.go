package main

import (
	"fmt"
	"os"

	"github.com/Virees/goru/flags"
	"github.com/Virees/goru/lib"
)

func main() {
	flags, err := flags.ParseInputFlags()
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}

	tc := lib.NewTagsCollector()

	urls := flags.GetUrls()
	for _, url := range urls {
		tc.Visit(url)
		if err != nil {
			fmt.Printf("Error visiting the site: %v", err)
		}
	}
}
