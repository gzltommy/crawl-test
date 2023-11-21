package scheduler

import (
	"github.com/gzltommy/crawl-test/04.MulTask/types"
)

type SimpleScheduler struct {
	workerChan chan types.Request
}

func NewSimpleScheduler() Scheduler {
	s := &SimpleScheduler{}
	return s
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan types.Request)
}

func (s *SimpleScheduler) WorkReady(r chan types.Request) {
	return
}

func (s *SimpleScheduler) WorkChan() chan types.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(request types.Request) {
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleScheduler) ConfigureWorkChan(c chan types.Request) {
	s.workerChan = c
}
