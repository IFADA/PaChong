package scheduler

import "PaChong/engine"

// the engine send the request to scheduler and scheduler distribute the request to different workers
type SimpleScheduler struct {
	workerChan chan engine.Request
}

//the ConfigureMasterWorkerChan will change the content of SimpleScheduler
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan  Submit(Request)
	//the program tend to get stuck
	//s.workerChan <- r
	go func() { s.workerChan <- r }()
}
