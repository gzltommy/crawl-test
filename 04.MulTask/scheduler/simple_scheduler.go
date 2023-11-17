package scheduler

import (
	"github.com/gzltommy/crawl-test/04.MulTask/types"
)

type Scheduler interface {
	Submit(request types.Request)
	ConfigureWorkChan(chan types.Request)
}

type SimpleScheduler struct {
	workerChan chan types.Request
}

func (s *SimpleScheduler) Submit(request types.Request) {
	s.workerChan <- request
}

func (s *SimpleScheduler) ConfigureWorkChan(c chan types.Request) {
	s.workerChan = c
}
