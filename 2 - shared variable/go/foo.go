package main

import (
	"fmt"
	"runtime"
)

func numberServer(add <-chan int, sub <-chan int, read chan<- int) {
	var number = 0

	// This for-select pattern is one you will become familiar with...
	for {
		select {
		case x := <-add:
			number += x
		case y := <-sub:
			number -= y
		case read <- number:
			// Do nothing
		}
	}
}

func incrementer(add chan<- int, finished chan<- struct{}) {
	for j := 0; j < 1000000; j++ {
		add <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- struct{}{}
}

func decrementer(sub chan<- int, finished chan<- struct{}) {
	for j := 0; j < 1000000+1; j++ {
		sub <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- struct{}{}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the remaining channels
	read := make(chan int) // Buffered channel with 1024 bytes
	add := make(chan int)
	sub := make(chan int)
	addFinished := make(chan struct{})
	subFinished := make(chan struct{})

	// TODO: Spawn the required goroutines
	go incrementer(add, addFinished)
	go decrementer(sub, subFinished)
	go numberServer(add, sub, read)

	// TODO: block on finished from both "worker" goroutines
	<-addFinished
	<-subFinished

	fmt.Println("The magic number is:", <-read)
}
