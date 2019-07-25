package processing

var creationQueue, workQueue chan int

func PrepareQueues() {
	processorCount := 5
	creationQueue = make(chan int)
	workQueue = make(chan int, processorCount)

	go processCreationQueue()

	for i := 0; i < processorCount; i++ {
		go processWorkQueue()
	}
}

func AddToQueue(id int) {
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