package engine

import (
	"PaChong/model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}
type Scheduler interface {
	ReadyNotifiter
	Submit(Request)
	// go to ask workchan
	WorkerChan() chan Request
	Run()
}
type ReadyNotifiter interface {
	WorkerReady(chan Request)
}

// the recipient of the pointer type
func (e *ConcurrentEngine) Run(seeds ...Request) {
	// all the worker user one input
	//in := make(chan Request)
	out := make(chan ParseResult)

	// ConfigureWorkerChan:the effect is put the Request in Scheduler
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}
	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got profile:#%d: %v", profileCount, item)
				profileCount++
			}

		}
		// Url dedup
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifiter) {
	go func() {
		for {
			//将Request chan send to Schedule的 workChan
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
