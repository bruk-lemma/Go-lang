package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	data, _ := ioutil.ReadFile("./notes.xml")

	note := &Notes{}

	_ = xml.Unmarshal([]byte(data), note)

	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Heading)
	fmt.Println(note.Body)
	write_xml_file()
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
