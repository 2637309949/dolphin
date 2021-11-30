package worker

import (
	"context"

	"github.com/2637309949/dolphin/packages/logrus"
)

// Example:
// // MyPayloadTest For example usage
// type MyPayloadTest struct {
// 	Num int
// }

// // Do do something
// func (p MyPayloadTest) Do() (err error) {
// 	//Do a job sample
// 	logrus.Info("I am working Do", p.Num)
// 	time.Sleep(time.Second * 2)
// 	err = nil
// 	return
// }

// func main() {
// 	dispatcher := NewDispatcher(10)
// 	dispatcher.Run()
//
// 	for i := 1; i < 40; i++ {
// 		job := NewJob(MyPayloadTest{i})
// 		dispatcher.AddJob(job)
// 	}
//
// 	time.Sleep(time.Second * 50)
// }

// Dispatcher -> Processing Data
type Dispatcher struct {
	// JobQueue A buffered channel that we can send work requests on.
	JobQueue chan Job
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	// Len represents num of worker
	Len int
}

// NewDispatcher represents Processing Data
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	jobQueue := make(chan Job, 1)
	return &Dispatcher{JobQueue: jobQueue, WorkerPool: pool, Len: maxWorkers}
}

// Run starting n number of workers
func (d *Dispatcher) Run() {
	for i := 0; i < d.Len; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			logrus.Info(context.TODO(), "Store a job into jobChannel")
			jobChannel := <-d.WorkerPool
			go func(job Job) {
				jobChannel <- job
			}(job)
		}
	}
}

// AddJob represents Generate Data
func (d *Dispatcher) AddJob(j Job) {
	d.JobQueue <- j
}

// Job defined
type Job interface {
	Do() (err error)
}

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// NewWorker represents the worker
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel
// in case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				if err := job.Do(); err != nil {
					logrus.Errorf(context.TODO(), "Error do payload function:%v", err.Error())
				}
			case <-w.quit:
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
