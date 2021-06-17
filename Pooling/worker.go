package pooling

import "log"

type Worker struct {
	ID               int //used to differentiate workers
	Dispatch_Channel chan chan *Job //Channel of channel of referenced jobs, shared between the workers and dispatcher,
	Worker_Channel   chan *Job //Channel of job reference, personal to each worker
	End              chan bool //Channel that signals an end to dispatcher
	Log              *log.Logger //logger instance to log results
}

func (w *Worker) Start() {
	go func() {
		for {
			w.Dispatch_Channel <- w.Worker_Channel //place the worker channel on the dispatch channel
			select {
				case task := <-w.Worker_Channel: //In the case that dispatcher send a job
					w.Log.Printf("Worker %d beginning task...", w.ID)
					err := task.Process() // process the job task
					w.Log.Printf("Worker %d has finished task: Awaiting new job...", w.ID)
					if err != nil { // if the task fails
						w.Log.Printf("Job Failed, sending End note now. Error : %s", err)
						w.End <- true // signal end to the dispatcher
					}
				case <-w.End:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.End <- true 
}