package queue

import (
	."Stringtheory-API/user-curriculum-progress/service-module"
	. "Stringtheory-API/user-curriculum-progress/types"
	"sync"
)

type QueueAdapter struct {
	workerPoolSize int
}

func NewQueueAdapter(workerPoolSize int) *QueueAdapter {
	return &QueueAdapter{ workerPoolSize: workerPoolSize }
}

func (adapter QueueAdapter) InitializeAdapter() {
	for worker := 1; worker <= adapter.workerPoolSize; worker++ {
		go adapter.spawnWorker()
	}
}

func (adapter QueueAdapter) spawnWorker() {
	for {
		message, err := UserCurriculumProgressModule.MessageStore.ReceiveMessage()
		if err != nil {
			// Skip to next iteration if there is no message
			continue
		}

		var waitGroup sync.WaitGroup
		waitGroup.Add(1)
			go func(message *Message) {
				defer waitGroup.Done()
			}(message)
		waitGroup.Wait()
	}
}
