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
					w.Log.Printf("Worker %d assigned task: %s ...processing now...", w.ID, task.Name)
					err := task.Process() // process the job task
					w.Log.Printf("Worker %d has finished task: %s,  Awaiting new job...", w.ID, task.Name)
					if err != nil { // if the task fails
						w.Log.Printf("Job Failed. Error : %s", err)
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