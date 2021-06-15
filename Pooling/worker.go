package pooling

type Work struct{
	ID int
	Job Job
}

type Worker struct{
	ID int
	Dispatch_Channel chan chan Work
	Worker_Channel chan Work
	End chan bool
}

func (w *Worker) Start(){
	go func(){
		for {
			w.Dispatch_Channel <- w.Worker_Channel
			select {
			case task := <- w.Worker_Channel:
				task.Job.Process()
			case <-w.End:
				return
			}
			
		}
	}()
}