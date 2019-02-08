package parser

import (
	"crawler/zhenai/engine"
	"regexp"
)

//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"`
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func ParseCity(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(content, -1)
	//log.Printf(">>>>>>>>\n %s ", matches)
	result := engine.ParseResult{}

	//log.Printf(">>>>>>>> %v /n",matches)
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParser,
			})
	}
	return result
}
