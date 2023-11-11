package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	go create_go_routine("https://www.golangprograms.com")
	go create_go_routine("https://coderwall.com")
	go create_go_routine("https://stackoverflow.com")
	time.Sleep(10 * time.Second)

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
