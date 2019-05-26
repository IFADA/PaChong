package scheduler

import "PaChong/engine"

type QueuedSheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedSheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedSheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (*QueuedSheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (s *QueuedSheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w //there is a worker ready to accept hte request
}

func (s *QueuedSheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	// submit and workerReady are two very separate things
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan: // Listen for read events
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
				// wo want request and workerchan are two separate program, so  if the request sended to workerchan,need delete all of them
			case activeWorker <- activeRequest: // Listen for writer events
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
