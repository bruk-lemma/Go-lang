package main

import (
	"fmt"
	"reflect"
)

func create_channels() {
	var ch chan int
	fmt.Println("The type of channels is", reflect.ValueOf(ch).Kind())
	//to send values to channel we use the <- operator
	//	ch <- 5 //send 5 into the channel
	fmt.Println("The vlaue of the channel ch is ", ch)
}

func create_buffred_channel() {
	/*ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 5*/

	//recieving valuse from a channel
	/*
		x := <-ch //x=3
		y := <-ch //y=5
	*/

}

func fibonacci_sequence_using_buffered_channels(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)

}

func doSomething(ch chan int) {
	x := 42
	ch <- x
}

func fibbonacci_sequence_using_unbuffred_channeals() {
	ch := make(chan int) //create an unbuffred channel;
	go doSomething(ch)   //launch a goroutine to do something with x

	x := <-ch //recieve x form the channel
	fmt.Println(x)
}

func worker(input chan int, output chan int, done chan bool) {
	for {
		select {
		case n := <-input:
			//do something on n
			output <- n * 2

		case <-done:
			close(output)
			return
		}
	}
}

func gracefull_shutdown_of_go_channels() {
	input := make(chan int)
	output := make(chan int)
	done := make(chan bool)

	go worker(input, output, done)

	// send some value on the input channel
	for i := 0; i < 10; i++ {
		input <- i
	}

	//close the input chaneel form the output channel

	close(input)

	//recieve values form the output channel

	for n := range output {
		fmt.Println(n)
	}

	//signal the worker to exit

	done <- true
}

func main() {
	// create_channels()
	// ch := make(chan int, 10)                              //create buffred channel with a acapcity of 10
	// go fibonacci_sequence_using_buffered_channels(10, ch) //generate the first 10 fibbonaci nu,bers in a separate goroutine

	// //read values form the channel until its closed
	// for x := range ch {
	// 	fmt.Println(x)
	// }

	//fibbonacci_sequence_using_unbuffred_channeals()
	gracefull_shutdown_of_go_channels()

}
