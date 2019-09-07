package javfree

import (
	"github.com/skyeidos/go-spider/src/spider/engine"
	"regexp"
)

var companyListRegex = regexp.MustCompile(`<a href="(https://javfree.me/category/mosaic/[^"]+)" data-instant="true">([^<]+)</a>`)

func CompanyListParser(content []byte, url string) engine.Result {
	matches := companyListRegex.FindAllSubmatch(content, -1)
	var requests []engine.Request
	for _, m := range matches {
		requests = append(requests, engine.Request{Url: string(m[1]), Parser: CompanyParser})
	}
	return engine.Result{Request: requests}
}
