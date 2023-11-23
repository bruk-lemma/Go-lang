package main

import (
	"fmt"
	"net"
)

func main() {

	//AA and AAAA
	iprecords, _ := net.LookupIP("facebook.com")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	//CNAME canonicalName

	cname, _ := net.LookupCNAME("m.facebook.com")
	fmt.Println(cname)

	//PTR pointer
	//These records provide the reverse binding from addresses to names
	ptr, _ := net.LookupAddr("6.8.8.8")
	for _, ptrvalue := range ptr {
		fmt.Println(ptrvalue)
	}

	//NS nameserver
	nameserver, _ := net.LookupNS("facebook.com")
	for _, ns := range nameserver {
		fmt.Println(ns)
	}
}
