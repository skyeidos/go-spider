package javfree

import (
	"github.com/gocolly/colly"
	"github.com/skyeidos/go-spider/src/spider/engine"
	"github.com/skyeidos/go-spider/src/spider/model"
	"log"
	"regexp"
)

var actorRegex = regexp.MustCompile(`出演者： ([^\s]+)`)
var releaseDateRegex = regexp.MustCompile(`発売日： ([^\s]+)`)
var durationRegex = regexp.MustCompile(`収録時間： ([^\s]+)`)
var idRegex = regexp.MustCompile(`https://javfree.me/(\d+)/[a-z0-9-]+`)

func MovieParser(_ []byte, title string, url string) engine.Result {
	c := getCollector()
	var info model.Movie
	var content string
	var images []string
	c.OnHTML("div.entry-content > p", func(element *colly.HTMLElement) {
		content = element.Text
	})
	c.OnHTML("div.entry-content > p > img", func(element *colly.HTMLElement) {
		images = append(images, element.Attr("src"))
	})
	if err := c.Visit(url); err != nil {
		log.Printf("MovieParser error: %v url:%s,", err, url)
	}
	id := parserString(idRegex, url)
	info.ReleaseDate = parserString(releaseDateRegex, content)
	info.Actor = parserString(actorRegex, content)
	info.Duration = parserString(durationRegex, content)
	info.Images = images
	info.Title = title
	var items []engine.Item
	items = append(items, engine.Item{Id: id, Payload: info})
	return engine.Result{
		Request: nil,
		Items:   items,
	}
}

func parserString(regex *regexp.Regexp, content string) string {
	bytes := []byte(content)
	matches := regex.FindAllSubmatch(bytes, -1)
	for _, m := range matches {
		return string(m[1])
	}
	return ""
}
