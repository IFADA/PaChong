package scheduler

import "PaChong/engine"

// the engine send the request to scheduler and scheduler distribute the request to different workers
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan  Submit(Request)
	//the program tend to get stuck
	//s.workerChan <- r
	go func() { s.workerChan <- r }()
}
