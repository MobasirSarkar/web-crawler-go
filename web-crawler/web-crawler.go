package webcrawler

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func WebCrawler() {
	c := colly.NewCollector()

	var url string
	fmt.Print(`Enter the Url : `)
	fmt.Scan(&url)

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		quote := h.ChildText("span.text")
		author := h.ChildText("small.author")
		tags := h.ChildText("div.tags")

		quote = strings.TrimSpace(quote)
		author = strings.TrimSpace(author)
		tags = strings.TrimSpace(tags)

		fmt.Printf("Quote :%s\nAuthor: %s\nTags: %s\n\n", quote, author, tags)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
