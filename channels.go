package main

import (
	"fmt"
	"reflect"
)

func main() {
	create_channels()
	ch := make(chan int, 10)                              //create buffred channel with a acapcity of 10
	go fibonacci_sequence_using_buffered_channels(10, ch) //generate the first 10 fibbonaci nu,bers in a separate goroutine

	//read values form the channel until its closed
	for x := range ch {
		fmt.Println(x)
	}

	fibbonacci_sequence_using_unbuffred_channeals()

}

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

func doSomrthing(ch chan int) {
	x := 42
	ch <- x
}

func fibbonacci_sequence_using_unbuffred_channeals() {
	ch := make(chan int) //create an unbuffred channel;
	go doSomrthing(ch)   //launch a goroutine to do something with x

	x := <-ch //recieve x form the channel
	fmt.Println(x)
}
