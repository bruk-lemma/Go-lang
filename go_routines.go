package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {

	// go create_go_routine("https://www.golangprograms.com")
	// go create_go_routine("https://coderwall.com")
	// go create_go_routine("https://stackoverflow.com")
	// time.Sleep(10 * time.Second)

	fmt.Println("Using WaitGroup")
	// Add a count of three, one for each goroutine.
	wg.Add(3)
	go using_WaitGroup_in_go_routines("https://www.golangprograms.com")
	go using_WaitGroup_in_go_routines("https://coderwall.com")
	go using_WaitGroup_in_go_routines("https://stackoverflow.com")

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Terminating Program")

}

func create_go_routine(url string) {
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step3 :", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step4 :", len(body))
}

var wg sync.WaitGroup

func using_WaitGroup_in_go_routines(url string) {
	//schedule the call to wait group;s Done to tell go routine  is complete
	defer wg.Done()

	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step 2:", url)
	defer response.Body.Close()

	fmt.Println("Step3 :", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4 :", len(body))

}
