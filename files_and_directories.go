package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// create_file()
	// create_directory()
	// rename_file()
	//copy_file_atspecified_location()
	//move_file_from_onelocation_toanother()
	metadata_of_file()

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

func rename_file() {
	oldName := "empty.txt"
	newName := "correctTest.txt"

	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

func copy_file_atspecified_location() {
	sourceFile, err := os.Open("/home/bruk/go/golang/testfolder/testfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer sourceFile.Close()

	//create new file
	newFile, err := os.Create("/home/bruk/go/golang/testfolder/testfolderin/testfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes. ", bytesCopied)

}

func move_file_from_onelocation_toanother() {
	newLocation := "/home/bruk/go/golang/newLocation/test.txt"
	oldLocation := "/home/bruk/go/golang/oldLocation/test.txt"

	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}

}

func metadata_of_file() {
	file, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileStat, err := os.Stat(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("FIle name:", fileStat.Name())
	fmt.Println("Size in bytes:", fileStat.Size())
	fmt.Println("Permissions:", fileStat.Mode())
	fmt.Println("Last modified:", fileStat.ModTime())
	fmt.Println("Is Directory:", fileStat.IsDir())

}
