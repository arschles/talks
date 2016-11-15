package main

// add is a convenience function for adder that makes it a blocking call
func add(i int, ch chan arg) int {
	retCh := make(chan int)
	ch <- arg{i: i, ret: retCh}
	retInt := <-retCh
	<-retCh // wait for the adder goroutine time finish
	return retInt
}
