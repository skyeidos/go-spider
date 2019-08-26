package javfree

import (
	"github.com/skyeidos/go-spider/src/spider/engine"
	"github.com/skyeidos/go-spider/src/spider/model"
	"regexp"
)

var actorRegex = regexp.MustCompile(`出演者： ([^<]+)<br/>`)
var releaseDateRegex = regexp.MustCompile("発売日： ([^<]+)<br/>")
var durationRegex = regexp.MustCompile(`収録時間： ([^<]+)<br/>`)
var imageRegex = regexp.MustCompile(`<img src="(https://cf.javfree.me/[^"]+).*?>`)
var idRegex = regexp.MustCompile(`https://javfree.me/(\d+)/[a-z0-9-]+`)

func MovieParser(content []byte, title string, url string) engine.Result {
	var info model.Movie
	id := parserString(idRegex, []byte(url))
	info.ReleaseDate = parserString(releaseDateRegex, content)
	info.Actor = parserString(actorRegex, content)
	info.Duration = parserString(durationRegex, content)
	info.Images = parserArray(imageRegex, content)
	info.Title = title
	var items []engine.Item
	items = append(items, engine.Item{Id: id, Payload: info})
	return engine.Result{
		Request: nil,
		Items:   items,
	}
}

func parserString(regex *regexp.Regexp, content []byte) string {
	matches := regex.FindAllSubmatch(content, -1)
	for _, m := range matches {
		return string(m[1])
	}
	return ""
}

func parserArray(regex *regexp.Regexp, content []byte) []string {
	var result []string
	matches := regex.FindAllSubmatch(content, -1)
	for _, m := range matches {
		result = append(result, string(m[1]))
	}
	return result
}
