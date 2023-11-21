package scheduler

import (
	"github.com/gzltommy/crawl-test/04.MulTask/types"
)

type QueueScheduler struct {
	requestChan chan types.Request
	workerChan  chan chan types.Request
}

func (s *QueueScheduler) WorkChan() chan types.Request {
	return make(chan types.Request)
}

func NewQueueScheduler() Scheduler {
	return &QueueScheduler{
		requestChan: make(chan types.Request),
		workerChan:  make(chan chan types.Request),
	}
}

func (s *QueueScheduler) Submit(r types.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) WorkReady(r chan types.Request) {
	s.workerChan <- r
}

func (s *QueueScheduler) Run() {
	go func() {
		var (
			requestQ []types.Request
			workQ    []chan types.Request
		)
		for {
			var (
				activeRequest types.Request
				activeWork    chan types.Request
			)
			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWork = workQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workQ = append(workQ, w)
			case activeWork <- activeRequest: // 将任务发送给空闲的工人
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
