package main

import (
	"PaChong/engine"
	"PaChong/scheduler"
	"PaChong/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedSheduler{}, WorkerCount: 100}

	e.Run(engine.Request{

		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//
	//	Url:        "https://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})

}
