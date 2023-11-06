//package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// )

type Config struct {
	SecretKey string `json:"secret_key"`
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error ocurred:%v\n", r)
			fmt.Println("Application terminated gracefully")
		} else {
			fmt.Println("Application executed successfully.")
		}
	}()
	//Read config file
	configData, err := ioutil.ReadFile("./config.json")
	if err != nil {
		// if there is an error reading the file,panic and terminate the program
		panic(fmt.Sprintf("Error reading config file: %s\n", err))
	}
	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		//If there is an error unmarshaling the JSON data,panic and terminate the program
		panic(fmt.Sprintf("Error unmarshaling JSON: %s\n", err))
	}
	if config.SecretKey == "" {
		//If the key is empty, panic and terminate the program
		panic("Sceret Key is missing form configuration")

	}
	//continue with the program execution
	fmt.Println("Application started successfully")
}
