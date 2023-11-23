package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	// data, _ := ioutil.ReadFile("./notes.xml")

	// note := &Notes{}

	// _ = xml.Unmarshal([]byte(data), note)

	// fmt.Println(note.To)
	// fmt.Println(note.From)
	// fmt.Println(note.Heading)
	// fmt.Println(note.Body)

	//write_xml_file()
	//read_json_file()
	write_json_file()
}

func write_xml_file() {

	note := Notes{To: "Nicky",
		From:    "Rock",
		Heading: "Meeting",
		Body:    "Meeting at 5pm!",
	}
	file, _ := xml.MarshalIndent(note, "", "")

	_ = ioutil.WriteFile("notes1.xml", file, 0644)
}

type CatalogNodes struct {
	CatalogNodes []Catalog `json:"catalog_nodes"`
}

type Catalog struct {
	Product_id string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

func read_json_file() {
	file, err := ioutil.ReadFile("catalog.json")
	if err != nil {
		log.Fatal(err)
	}
	data := CatalogNodes{}

	_ = json.Unmarshal([]byte(file), &data)
	fmt.Println("Reading from json file", data)
	for i := 0; i < len(data.CatalogNodes); i++ {
		fmt.Println("Product Id:", data.CatalogNodes[i].Product_id)
		fmt.Println("Quantity: ", data.CatalogNodes[i].Quantity)
	}

}

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func write_json_file() {
	data := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}
