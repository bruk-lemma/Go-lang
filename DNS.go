// package main

// import (
// 	"fmt"
// 	"net"
// )

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

	//MX records
	//These records identify the servers that can exchange emails

	mxrecords, _ := net.LookupMX("facebook.com")
	for _, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}

	//SRV service record
	/*
		The LookupSRV function tries to resolve an SRV query of the given service,
		protocol, and domain name. The second parameter is "tcp" or "udp"*/
	cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", "facebook.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n cname: %s \n\n", cname)

	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}

	//TXT record
	/*
		This text record
		stores information about the SPF that
		can identify the authorized server to send email on behalf of your organization*/

	txtrecords, _ := net.LookupTXT("facebook.com")
	for _, txt := range txtrecords {
		fmt.Println(txt)
	}

}
