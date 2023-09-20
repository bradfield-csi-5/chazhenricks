package main

import (
	"fmt"
	"net"
	"os"
)

const (
	GoogleDns = "8.8.8.8"
	DnsPort   = 53
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//take in URL as user input
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("need to provide args")
	}
	fmt.Printf("%s\n", args[0])

	//construct a DNS message and open a socket to the DNS server
	hostPort := fmt.Sprintf("%s:%d", GoogleDns, DnsPort)
	udpAddr, err := net.ResolveUDPAddr("udp4", hostPort)
	checkErr(err)

	connection, err := net.DialUDP("udp4", nil, udpAddr)
	checkErr(err)

  fmt.Printf("Udp Server: %s\n", connection.RemoteAddr().String())

}
