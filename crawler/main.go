package main

import (
	"crawler/scheduler"
	"crawler/zhenai/engine"
	"crawler/zhenai/parser"
	"fmt"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	fmt.Printf("%d,\n", 22222222)
}
