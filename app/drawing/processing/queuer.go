package processing

var creationQueue, workQueue chan int64

func PrepareQueues() {
	processorCount := 5
	creationQueue = make(chan int64)
	workQueue = make(chan int64, processorCount)

	go processCreationQueue()

	for i := 0; i < processorCount; i++ {
		go processWorkQueue()
	}
}

func AddToQueue(id int64) {
	creationQueue <- id
}

func processCreationQueue() {
	for {
		select {
		case id := <-creationQueue:
			workQueue <- id
		}
	}
}

func processWorkQueue() {
	for {
		select {
		case id := <-workQueue:
			Process(id)
		}
	}
}
