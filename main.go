package main

import (
	"fmt"
	"os"

	"github.com/Virees/goru/lib"
)

func main() {
	flags, err := lib.ParseInputFlags()
	if err != nil {
		fmt.Println("Error parsing flags: ", err)
		os.Exit(1)
	}

	tc := lib.NewTagsCollector()

	url, err := flags.GetUrl()
	if err != nil {
		fmt.Println("Error generating URL: ", err)
		os.Exit(1)
	}

	err = tc.Visit(url)
	if err != nil {
		fmt.Println("Error visiting the site: ", err)
		os.Exit(1)
	}
}
