package main

import (
	"log"
	"os"
)

func main() {
	create_file()
	create_directory()

}

func create_file() {
	//create file
	emptyFile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)

	}
	log.Println(emptyFile)
	emptyFile.Close()
}

func create_directory() {
	//create directory.
	_, err := os.Stat("test")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}
