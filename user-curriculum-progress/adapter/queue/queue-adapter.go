package queue

import (
	."Stringtheory-API/user-curriculum-progress/service-module"
	"sync"
)

type QueueAdapter struct {
	workerPoolSize int
}

func NewQueueAdapter(workerPoolSize int, queueURI string) *QueueAdapter {
	return &QueueAdapter{ workerPoolSize: workerPoolSize }
}

func (adapter QueueAdapter) InitializeAdapter() {
	for worker := 1; worker <= adapter.workerPoolSize; worker++ {
		go adapter.spawnWorker()
	}
}

func (adapter QueueAdapter) spawnWorker() {
	for {
		messages, err := UserCurriculumProgressModule.MessageStore.ReceiveMessages()
		if err != nil {
			// Skip to next iteration if there is no message
			continue
		}

		var waitGroup sync.WaitGroup
		for _, message := range messages {
			waitGroup.Add(1)
			go func(m string) {
				defer waitGroup.Done()
			}(message)
			waitGroup.Wait()
		}
	}
}
