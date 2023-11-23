// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"strings"
// )

func main() {
	// create_file()
	// create_directory()
	// rename_file()
	//copy_file_atspecified_location()
	//move_file_from_onelocation_toanother()
	//metadata_of_file()
	//delete_file()
	readfile_character_by_character("buffertext.txt")
	truncate_file()

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

func delete_file() {
	err := os.Remove("empty.txt")

	if err != nil {
		log.Fatal(err)
	}

}

func readfile_character_by_character(filename string) {
	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputData := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputData))
	data.Split(bufio.ScanRunes)

	for data.Scan() {
		fmt.Print(data.Text())
	}

}

func truncate_file() {
	//truncate if file size exceeds 100kb
	err := os.Truncate("correctTest.txt", 100)
	if err != nil {
		log.Fatal(err)
	}

}

func add_content_to_file() {
	message := "Addt this content at the end"
	filename := "test.txt"

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)

	}
	defer f.Close()
	fmt.Fprintf(f, "%s\n", message)
}
