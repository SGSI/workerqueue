package worker

import (
	"fmt"
)

type Worker struct {
	WorkerID      *string
	WorkerName    *string
	WorkerChannel chan interface{}
}

func NewWorker(workerID string, workerName string) Worker {
	w := Worker{
		WorkerID:      &workerID,
		WorkerName:    &workerName,
		WorkerChannel: make(chan interface{}, 1),
	}
	return w
}

func CreateWorkers(NoOfWorkers int, workerName string) []Worker {
	workers := []Worker{}
	for id := 0; id < NoOfWorkers; id++ {
		workers = append(workers, NewWorker(fmt.Sprintf("%d", id), fmt.Sprintf("%s-%d", workerName, id)))
	}
	return workers
}
