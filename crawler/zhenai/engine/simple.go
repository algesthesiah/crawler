package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("got 1url %s", r.Url)
		parseResult, err := worker(r)
		if err != nil {
			log.Printf("Fetch :error fetching url %s %v", r.Url, err)
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("got item %s", item)
		}
	}
}
func worker(
	r Request) (ParseResult, error) {
		log.Printf("fetching %s",r.Url)
	body, e := fetcher.Fetch(r.Url)
	if e != nil {
		log.Printf("FeatchErr:", e)
		return ParseResult{},e
	}
	return r.ParserFunc(body), nil
}
