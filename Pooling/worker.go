package pooling

type Worker struct{
	ID int
	Dispatch_Channel chan chan *Job
	Worker_Channel chan *Job
	End chan bool
}

func (w *Worker) Start(){
	go func(){
		for {
			w.Dispatch_Channel <- w.Worker_Channel
			select {
				case task := <- w.Worker_Channel:
					task.Process()
				case <-w.End:
					return
			}
		}
	}()
}

func (w *Worker) Stop(){
	w.End <- true
}