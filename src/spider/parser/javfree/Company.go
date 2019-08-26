package javfree

import (
	"github.com/skyeidos/go-spider/src/spider/engine"
	"regexp"
)

var companyRegex = regexp.MustCompile(`<h2 class="entry-title"><a href="(https://javfree.me/[^"]+)">([^<]+)</a></h2>`)
var nextPageRegex = regexp.MustCompile(`<a class="next page-numbers" href="(https://javfree.me/category/mosaic/prestige-mosaic/page/[0-9]+)">Next</a>`)

func CompanyParser(content []byte) engine.Result {
	matches := companyRegex.FindAllSubmatch(content, -1)
	var requests []engine.Request
	for _, m := range matches {
		name := string(m[2])
		url := string(m[1])
		requests = append(requests, engine.Request{Url: url, Parser: func(content []byte) engine.Result {
			return MovieParser(content, name, url)
		}})
	}
	matches = nextPageRegex.FindAllSubmatch(content, -1)
	for _, m := range matches {
		requests = append(requests, engine.Request{Url: string(m[1]), Parser: CompanyParser})
	}
	return engine.Result{Request: requests}
}
