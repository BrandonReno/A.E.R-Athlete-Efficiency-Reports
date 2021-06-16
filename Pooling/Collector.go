package pooling

var AvailableWorkers = make(chan chan *Job)

type Collector struct {
	Work chan Job //Jobs come in through this channel
	End chan bool //End notice comes in through this channel
}

func StartDispatcher(workerCount int) Collector {
	// Create the collector

	incomingJobs := make(chan Job)
	endWork := make(chan bool)
	collector := Collector{Work: incomingJobs, End: endWork}
	var workers []Worker

	//Initialize the workers

	for i := 0; i < workerCount; i++{
		w := Worker{
			ID: i,
			Dispatch_Channel: AvailableWorkers,
			Worker_Channel: make(chan *Job),
			End: make(chan bool),
		}
		workers = append(workers, w)

		w.Start()
	}

	go func(){
		for {
			select{
				case <-endWork:
					for _,w := range workers{
						w.Stop()
					}
					return
				case job := <-incomingJobs:
					nextAvailableWorker := <- AvailableWorkers
					nextAvailableWorker <- &job
			}

				

		}
	}()
	
	return collector
}