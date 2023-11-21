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
	<-done
}

func recieve_data_from_multiple_channels_using_for_slecet() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch1:
			if ok {
				fmt.Println("Recieved from ch1:", x)
			} else {
				fmt.Println("ch1 closed")
				//close(ch1)
			}
		case x, ok := <-ch2:
			if ok {
				fmt.Println("Recieved from ch2:", x)
			} else {
				fmt.Println("ch2 close")

			}
		}
	}
}

func recieve_from_closed_channel() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for {
		i, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("recieved", i)
	}
}

func accept_channel_as_parameter(ch chan<- int) {
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
}

func use_func_that_accepts_chan_as_param() {
	ch := make(chan int)
	go accept_channel_as_parameter(ch)
	for i := range ch {
		fmt.Println(i)
	}
}

func broadCastMessage_using_channel() {
	ch := make(chan string, 1)

	//start three goroutines taat read from the channel
	for i := 0; i < 3; i++ {
		go func() {
			msg := <-ch
			fmt.Println("Recieved message:", msg)
		}()

		//Broadcast a message to all the goroutines
		ch <- "hello,world"
	}

}

func howtoiterate_over_channels() {
	//using for range loop
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	for value := range ch {
		fmt.Println(value)
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case val := <-ch1:
			fmt.Println("Recieved from ch1", val)
		case val := <-ch2:
			fmt.Println("Recieved from ch2", val)
		}
	}

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
	//gracefull_shutdown_of_go_channels()

	//recieve_data_from_multiple_channels_using_for_slecet()
	// recieve_from_closed_channel()
	// use_func_that_accepts_chan_as_param()
	// broadCastMessage_using_channel()
	howtoiterate_over_channels()

}
