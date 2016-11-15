package main

// start up workers, each waiting for a submission
func startWorkers(ch chan workItem, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go wrk(ch)
	}
}
