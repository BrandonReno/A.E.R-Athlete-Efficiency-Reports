package pooling

import "log"

var AvailableWorkers = make(chan chan *Job)

type Collector struct {
	Work chan Job //Jobs come in through this channel
	End chan bool //End notice comes in through this channel to shut everything down
}

func StartDispatcher(workerCount int, l *log.Logger) *Collector {
	// First Create the collector which collects jobs

	incomingJobs := make(chan Job) // Channel of jobs that come in to the collector
	endWork := make(chan bool) // endwork channel that shuts all the workers down if anything is sent through
	collector := Collector{Work: incomingJobs, End: endWork}
	var workers []Worker

	//Initialize the workers

	for i := 0; i < workerCount; i++{
		w := Worker{
			ID: i,
			Dispatch_Channel: AvailableWorkers,
			Worker_Channel: make(chan *Job),
			End: make(chan bool),
			Log: l,
		}
		
		workers = append(workers, w)

		w.Start() //start all the workers
	}

	// Start the dispatcher
	go func(){
		l.Println("Listening for Jobs...")
		for {
			select{
				case <-endWork:
					for _,w := range workers{
						l.Println("Worker channels closing.")
						w.Stop()
					}
					return
				case job := <-incomingJobs: //when an incoming job comes in
					nextAvailableWorker := <- AvailableWorkers // check who the next available worker in the Available worker channel
					nextAvailableWorker <- &job // send that next available worker the job to process
			}

				

		}
	}()
	
	return &collector
}

func (c *Collector) EnqueJob(j *Job){
	c.Work <- *j
}

func (c *Collector) EndProcesses(){
	c.End <- true
}