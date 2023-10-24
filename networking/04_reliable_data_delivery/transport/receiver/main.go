package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"golang.org/x/sys/unix"
)

type Header struct {
	Awk      int
	Seq      int
	Checksum byte
}

type Message struct {
	Header Header
	Data   []byte
}

func MakeChecksum(data []byte) byte {
	var checksum byte
	for _, b := range data {
		checksum ^= b
	}
	return checksum
}
func isCurrupt(data []byte, receivedChecksum byte) bool {
	checksum := MakeChecksum(data)
	return checksum == receivedChecksum
}

func isSeq0(seq int) bool {
	return seq == 0
}

func isSeq1(seq int) bool {
	return seq == 1
}

func CreatePacket(awk int, seq int, checksum byte, data []byte) *Message {
	return &Message{
		Header: Header{
			Awk:      awk,
			Seq:      seq,
			Checksum: checksum,
		},
		Data: data,
	}
}

func main() {

	args := os.Args
	if len(args) != 2 {
		panic("yah need a port ya dingus")
	}
	port := os.Args[1]

	//create listen socket
	ln, err := net.Listen("udp", fmt.Sprintf(":%s", port))
	defer ln.Close()
	check(err)

	//listen for connections
	fmt.Printf("Server listening on port: %d\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error accepting connection", err)
			continue
		}
		handleConnection(conn)

	}
}

func handleConnection(conn net.Conn){
  buf := make([]byte, 0, 1024) 
  _, err := conn.Read(buf)

  if err != nil {
    if err != io.EOF{
      check(err)
    }
  }
  fmt.Println("total size:", len(buf))
}

// func CreateSocket(host string, port int, action string) int {
//
// 	socketFd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
// 	check(err)
//
// 	// Resolve the IP address to a numeric address
// 	ip := net.ParseIP(host)
// 	ipBytes := ip.To4()
// 	if ipBytes == nil {
// 		fmt.Println("Invalid IP address:", host)
// 		os.Exit(1)
// 	}
//
// 	// Create the sockaddr structure
// 	sockaddr := unix.SockaddrInet4{
// 		Port: port,
// 	}
// 	copy(sockaddr.Addr[:], ipBytes)
//
// 	// Bind the socket
// 	switch action {
// 	case "connect":
// 		err := unix.Connect(socketFd, &sockaddr)
// 		check(err)
// 	case "bind":
// 		err := unix.Bind(socketFd, &sockaddr)
// 		check(err)
// 	}
// 	return socketFd
// }
//
// func receivePacket(fd int, buf *[]byte){
//   bytesRead, _, err := unix.Recvfrom(fd, *buf, 0)
//   if err != nil {
//     fmt.Println("err reading from Recvfrom", err)
//   }
//   if bytesRead == 0{
//     fmt.Println("Conncection closed - nothing received")
//   }
// }

func check(err error) {
	if err != nil {
		panic(err)
	}
}
