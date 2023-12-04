package main

import (
	"fmt"

	"github.com/Virees/goru/goru"
	"github.com/gocolly/colly/v2"
)

func setupColly(c *colly.Collector) {
	c.OnHTML(".tag", func(e *colly.HTMLElement) {
		tags, err := goru.ParseBooruTags(e.Text)
		if err != nil {
			fmt.Println("Tags parse error: ", err)
		}

		fmt.Printf("%s\n", tags)
	})
}

func main() {
	id := goru.ParseInputFlags()
	c := colly.NewCollector()

	setupColly(c)

	err := c.Visit("" + fmt.Sprintf("https://safebooru.org/index.php?page=post&s=view&id=%d", id))
	if err != nil {
		fmt.Println("Error visiting the site:", err)
	}
}
