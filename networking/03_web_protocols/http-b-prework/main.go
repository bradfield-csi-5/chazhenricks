package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

var globalCache = make(map[string][]byte)

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

func forward(clientFd, serverFd int, action string) {
	fmt.Printf("ACTION:%s\n", action)
	defer unix.Close(clientFd)
	defer unix.Close(serverFd)

	buffer := make([]byte, 1024)
	for {
		// Read data from the client
		fmt.Printf("I AM READING. ClientFd: %d\n", clientFd)
		bytesRead, err := unix.Read(clientFd, buffer)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}

		if bytesRead == 0 {
			// Connection closed by the client
			return
		}

		_, err = unix.Write(serverFd, buffer[:bytesRead])
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
	}
}

func handleConnection(clientFd int) {

	buffer := make([]byte, 1024)
	bytesRead, err := unix.Read(clientFd, buffer)
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}

	request := string(buffer[:bytesRead])
	lines := strings.Split(request, "\r\n")

	if len(lines) < 1 {
		fmt.Println("Invalid http request - no lines read in request")
		unix.Close(clientFd)
		return
	}

	firstLine := strings.Fields(lines[0])
	if len(firstLine) < 3 {
		fmt.Println("Invalid http request not enough items in first line header")
		unix.Close(clientFd)
		return
	}

	// method := firstLine[0]
	path := firstLine[1]

	value, exists := globalCache[path]
	if exists {
		fmt.Printf("LOL I HIT THE CACHE")
		// hit cache, write cached value to client and return
		_, err = unix.Write(clientFd, value)
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
		unix.Close(clientFd)
	} else {
		fmt.Println("I DIDNT HIT THECACHE")
		secondLine := strings.Fields(lines[1])
		clientHost := secondLine[1]
		parts := strings.Split(clientHost, ":")
		port, err := strconv.Atoi(parts[1])
		check(err)
		newClientFd := CreateSocket(parts[0], port, "connect")
		serverFd := CreateSocket("127.0.0.1", 9000, "connect")
		go forward(newClientFd, serverFd, "clientSending")
		go forward(serverFd, newClientFd, "clientReceiving")
	}

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
