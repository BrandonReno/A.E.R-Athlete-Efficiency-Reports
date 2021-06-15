package pooling

type Work struct{
	ID int
	Job Job
}

type Worker struct{
	ID int
	Dispatch_Channel chan Work
}