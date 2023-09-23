package main

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/sys/unix"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateSocket(host string, port int, action string) int {

	socketFd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		os.Exit(1)
	}

	// Define the IP address and port

	// Resolve the IP address to a numeric address
	ip := net.ParseIP(host)
	ipBytes := ip.To4()
	if ipBytes == nil {
		fmt.Println("Invalid IP address:", host)
		os.Exit(1)
	}

	// Create the sockaddr structure
	sockaddr := unix.SockaddrInet4{
		Port: port,
	}
	copy(sockaddr.Addr[:], ipBytes)

	// Bind the socket
	switch action {
	case "connect":
		fmt.Println("I connected")
		err := unix.Connect(socketFd, &sockaddr)
		check(err)
	case "bind":
		fmt.Println("I BINDED")
		err := unix.Bind(socketFd, &sockaddr)
		check(err)
	}

	return socketFd
}

func main() {
	// Create a TCP socket
	ipAddress := "127.0.0.1"
	port := 8080
	listenSocketFd := CreateSocket(ipAddress, port, "bind")

	defer unix.Close(listenSocketFd)

	// Listen for incoming connections
	if err := unix.Listen(listenSocketFd, 5); err != nil {
		fmt.Println("Error listening on socket:", err)
		os.Exit(1)
	}
	fmt.Println("Server is listening on", ipAddress+":", port)

	for {
		// Accept incoming connections
		clientFd, _, err := unix.Accept(listenSocketFd)
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("I GOT A CONNECTION")

		go handleConnection(clientFd)
	}
}

func handleConnection(clientFd int) {
	defer unix.Close(clientFd)
	fmt.Println("WHAT")
	buffer := make([]byte, 1024)
	for {
		// Read data from the client
		bytesRead, err := unix.Read(clientFd, buffer)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}

		if bytesRead == 0 {
			// Connection closed by the client
			return
		}

		destFd := CreateSocket("127.0.0.1", 5678, "connect")
		// Echo the data back to the client
		_, err = unix.Write(destFd, buffer[:bytesRead])
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
	}
}
