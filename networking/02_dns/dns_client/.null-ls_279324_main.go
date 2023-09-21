package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	GoogleDns = "8.8.8.8"
	DnsPort   = 53
)

// BinaryTypes
var big = binary.BigEndian
var little = binary.LittleEndian

// RR Types
type Type uint16

const (
	A     Type = 1 //host address
	NS         = 2 //authoritative name server
	MD         = 3
	MF         = 4
	CNAME      = 5 //canonical name for an alias
	SOA        = 6
	MB         = 7
	MG         = 8
	MR         = 9
	NULL       = 10
	WKS        = 11
	PTR        = 12
	HINFO      = 13
	MINFO      = 14
	MX         = 15 //mail server
	TXT        = 16
)

// Resource Record Class Values
type Class uint16

const (
	IN Class = 1 //the internet
	CS       = 2 //the CSNet
	CH       = 3 //CHAOS
	HS       = 4
)

type DnsHeader struct {
	Id      uint16
	Flags   uint16
	QdCount uint16
	AnCount uint16
	NsCount uint16
	ArCount uint16
}

type Question struct {
	QName  []byte
	QType  Type
	QClass Class
}

type ResourceRecord struct {
	Name     string
	Type     Type
	Class    Class
	TTL      uint32
	RDLength uint16
	RData    string
}

type Message struct {
	Header     DnsHeader
	Question   Question
	Answer     uint32
	Authority  uint32
	Additional uint32
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func BuildHeader(buf *bytes.Buffer) {
	header := DnsHeader{
		66,
		0x0120,
		1,
		0,
		0,
		0,
	}

	binary.Write(buf, big, header.Id)
	binary.Write(buf, big, header.Flags)
	binary.Write(buf, big, header.QdCount)
	binary.Write(buf, big, header.AnCount)
	binary.Write(buf, big, header.ArCount)
	binary.Write(buf, big, header.NsCount)
}

func BuildQuestion(buf *bytes.Buffer, name string) {
	question := Question{
		BuildQName(name),
		A,
		IN,
	}

	binary.Write(buf, big, question.QName)
	binary.Write(buf, big, question.QType)
	binary.Write(buf, big, question.QClass)
}

func BuildQName(name string) []byte {

	var buf bytes.Buffer
	split := strings.Split(name, ".")
	for _, value := range split {

		length := byte(len(value))
		lengthSlice := []byte{length}
		_, err := buf.Write(lengthSlice)
		checkErr(err)
		_, err = buf.Write([]byte(value))
		checkErr(err)
	}

	return buf.Bytes()
}

func main() {
	//take in URL as user input
	arg := os.Args[1]
	fmt.Printf("%s\n", arg)

	conn, err := net.ListenPacket("udp", ":0")
	checkErr(err)
	defer conn.Close()

	dnsServer := fmt.Sprintf("%s:%d", GoogleDns, DnsPort)
	dst, err := net.ResolveUDPAddr("udp", dnsServer)
	checkErr(err)

	//build message
	//start with byte buffer and pass it to the different building blocks
	var buf bytes.Buffer
	BuildHeader(&buf)
	BuildQuestion(&buf, arg)

	message := buf.Bytes()
	fmt.Printf("MESSAGE %v\n", message)

	_, err = conn.WriteTo(message, dst)

	res := bytes.NewBuffer(make([]byte, 1024))
	conn.ReadFrom(res.Bytes())
	fmt.Printf("WHAT DID I GET? %v\n", res)
	// fmt.Printf("WHATS THE STATUS %s\n", status)
}
