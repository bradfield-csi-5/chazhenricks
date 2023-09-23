// package main
//
// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"net"
// )
//
// func check(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
// func main() {
// 	fmt.Println("Hello")
//
// 	listener, err := net.Listen("tcp", "127.0.0.1:8080")
// 	check(err)
// 	defer listener.Close()
//
// 	for {
// 		conn, err := listener.Accept()
// 		check(err)
// 		go handleRequests(conn)
// 	}
// }
//
// func handleRequests(conn net.Conn) {
// 	var buf bytes.Buffer
// 	_, err := io.Copy(&buf, conn)
// 	check(err)
//
// 	fmt.Printf("I WROTE SOME STUFF %s\n", buf.String())
// }
