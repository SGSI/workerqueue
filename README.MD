# WorkerQueue

WorkerQueue is a package to create P workers to perform Q tasks from the Queue of length R parallelly.

## Install

```sh
go get -u github.com/SGSI/workerqueue
```

## Example

```go
import (
	"fmt"
	"github.com/SGSI/workerqueue"
	"time"
)
const (
	QueueLength = 1000 // Length of the Queue
	NoOfWorkers = 5    // No of Workers
)

// This is a worker function. This function will be executed by the worker with provided data
func workerFunction(params ...interface{}) {
	fmt.Println("working with data = ", params)
}

// This is a dummy function to add data to queue
func AddToQueue(queue workerqueue.Queue) {
	idx := 0
    topic := "Topic"
	for {
		idx += 1
		queueData := workerqueue.QueueData{
			Data:      idx,
			Topic:     &topic,
			TimeStamp: time.Now(),
			MetaData: map[string]interface{}{
				"key": "value",
			},
		}
        // Enqueue the data to queue
		queue.Enqueue(queueData)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	// creating queue.
	queue := workerqueue.CreateQueue()

	// initialising queue.
	err := queue.Init("QueueType", "QueueName", QueueLength, NoOfWorkers, workerFunction)
	if err != nil {
		panic(err)
	}

	// function to add data to queue
	go AddToQueue(queue)
	time.Sleep(1 * time.Hour)
}
```
