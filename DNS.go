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

}
