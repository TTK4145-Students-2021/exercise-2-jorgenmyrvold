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
		case x := <-sub:
			number -= x
		}
		// read <- number
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
	read := make(chan int, 1024) // Buffered channel with 1024 bytes
	add := make(chan int, 1024)
	sub := make(chan int, 1024)
	addFinished := make(chan struct{})
	subFinished := make(chan struct{})

	// TODO: Spawn the required goroutines
	go incrementer(add, addFinished)
	go decrementer(sub, subFinished)
	go numberServer(add, sub, read)

	// time.Sleep(time.Millisecond * 1000)
	fmt.Println("The magic number is:", <-read)

	// TODO: block on finished from both "worker" goroutines
	dummy1 := <-addFinished // Not possible to write _ := <-addFinished
	dummy2 := <-subFinished
	fmt.Println(dummy1, dummy2) // Must use variables because not possible to write as comment on line 54

}
