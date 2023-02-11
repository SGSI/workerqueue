package workerqueue

import (
	"fmt"
	"github.com/SGSI/workerqueue/worker"
	"sync"
	"time"
)

var dequeueMutex = &sync.Mutex{}
var enqueueMutex = &sync.Mutex{}

type QueueData struct {
	Data      interface{}
	Topic     *string
	TimeStamp time.Time
	MetaData  map[string]interface{}
}

type Queue struct {
	Type        *string
	Name        *string
	Data        chan interface{}
	Size        *int
	NoOfWorkers *int
	WorkerFunc  func(...interface{})
	Workers     []worker.Worker
}

func (q *Queue) Enqueue(data interface{}) error {
	enqueueMutex.Lock()
	q.Data <- data
	*q.Size += 1
	enqueueMutex.Unlock()
	return nil
}

func (q *Queue) Dequeue() {
	dequeueMutex.Lock()
	*q.Size -= 1
	dequeueMutex.Unlock()
}

func (q *Queue) Length() *int {
	return q.Size
}
func CreateQueue() Queue {
	newQueue := Queue{}
	return newQueue
}
func (q *Queue) Init(queueType string, QueueName string, QueueSize int, NoOfWorkers int, workerFunc func(...interface{})) error {
	fmt.Println("Initialising Queue")
	q.Type = &queueType
	q.Name = &QueueName
	q.Size = &QueueSize
	q.NoOfWorkers = &NoOfWorkers
	q.Data = make(chan interface{}, QueueSize)
	q.WorkerFunc = workerFunc
	q.Workers = worker.CreateWorkers(*q.NoOfWorkers, *q.Name)
	go q.startWorkers()
	go q.startProvidingWorkToWorkers()
	return nil
}

func (q *Queue) startProvidingWorkToWorkers() {
	for _, wrker := range q.Workers {
		go func(w worker.Worker) {
			for {
				incomingWork := <-q.Data
				w.WorkerChannel <- incomingWork
			}
		}(wrker)
	}
}

func (q *Queue) startWorkers() {
	for _, wrkr := range q.Workers {
		go func(w worker.Worker) {
			for {
				workData := <-w.WorkerChannel
				q.WorkerFunc(workData)
				q.Dequeue()
			}
		}(wrkr)
	}
}
