package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"unsafe"
)

const (
	GlobalHeaderLength = 24
	PerPacketHeaderLength   = 16
)

type GlobalHeader struct {
	Magic               uint32
	MajorV              uint16
	MinorV              uint16
	TimeZoneOffset      uint32
	Accuracy            uint32
	SnapshotLength      uint32
	LinkLayerHeaderType uint32
}

type PerPacketHeader struct {
	TimeStampSeconds  uint32
	NanoMicro         uint32
	PacketLength      uint32
	UnTruncatedLength uint32
}

type EthernetFrames struct {
	source uint64
	data   []byte
}

type DecodedPacket struct {
	order int
	data  []byte
}
type DecodedImage struct {
	packets []DecodedPacket
}

func main() {
	fmt.Println("hello")
	// get data in byte array
	data, err := LoadData()
	if err != nil {
		fmt.Errorf("error in load data %v\n", err)
	}

	//parse into global header

	header := ParseGlobalHeader(data)
	fmt.Printf("Global Header Data: %+v\n", header)
	fmt.Printf("Size of Global Header: %d \n", unsafe.Sizeof(header))
	packets := CountPackets(data[24:])
	fmt.Printf("Number of Packets: %d\n", len(packets))

	ethernetFrames := ParseEthernetFrame(packets)
	ipDatagrams := ParseIpDatagram(ethernetFrames)
	tcpSegments := ParseTCPSegment(ipDatagrams)

	// fmt.Printf("TCP PACKETS: %d \n", len(tcpSegments.packets))

	CreateImage(tcpSegments)
}

func LoadData() (capture []byte, error error) {
	file, err := ioutil.ReadFile("net.cap")
	if err != nil {
		return nil, fmt.Errorf("unable to read net.cap %v", err)
	}

	buffer := make([]byte, len(file))
	copy(buffer, file)
	return buffer, nil

}

func ParseGlobalHeader(data []byte) GlobalHeader {
	header := GlobalHeader{}
	header.Magic = binary.LittleEndian.Uint32(data[:4])
	header.MajorV = binary.LittleEndian.Uint16(data[4:6])
	header.MinorV = binary.LittleEndian.Uint16(data[6:8])
	header.TimeZoneOffset = binary.LittleEndian.Uint32(data[8:12])
	header.Accuracy = binary.LittleEndian.Uint32(data[12:16])
	header.SnapshotLength = binary.LittleEndian.Uint32(data[16:20])
	header.LinkLayerHeaderType = binary.LittleEndian.Uint32(data[20:24])

	return header
}

func CountPackets(data []byte) [][]byte {
	packets := make([][]byte, 0)
	var packetIndex uint32 = 0
	var dataLength uint32 = uint32(len(data))

	for packetIndex < dataLength {

		//size is a 4 byte number 8 bytes from the start of the header
		sizeIndex := packetIndex + 8
		packetSize := binary.LittleEndian.Uint32(data[sizeIndex : sizeIndex+4])
		fmt.Printf("PACKET SIZE: %d\n", packetSize)

		endOfPacket := packetIndex + 16 + packetSize
		packetPayloadIndex := packetIndex + 16
		packetIndex = endOfPacket
		packet := data[packetPayloadIndex:endOfPacket]
		packets = append(packets, packet)
	}

	return packets
}

func ParseEthernetFrame(packets [][]byte) [][]byte {

	frames := make([][]byte, 0)

	for _, packet := range packets {

		var destMacAddress uint64 = Parse6ByteNumber(packet[:6])
		var sourceMacAddress uint64 = Parse6ByteNumber(packet[6:12])
		ipType := DecodeIPType(packet[12:14])
		fmt.Println("---Ethernet Frame---")
		fmt.Printf("DEST MAC: %x - SOURCE MAC %x - IP VERSION: %s\n", destMacAddress, sourceMacAddress, ipType)

		// endIndex := len(packet) - 4
		frame := packet[14:]
		frames = append(frames, frame)
	}
	return frames

}

func ParseIpDatagram(frames [][]byte) [][]byte {
	datagrams := make([][]byte, 0)

	for _, frame := range frames {
		headerLength := Parse4BitNumber(frame[0])
		headerSize := (headerLength * 32) / 8
		lengthOfPayload := len(frame) - headerSize
		fmt.Printf("SIZE OF DATAGRAM: %d - HEADER: %d - PAYLOAD:%d\n", len(frame), headerSize, lengthOfPayload)

		fmt.Printf("TRANSPORT PROTOCOL: %d\n", frame[9])

		datagram := frame[headerSize:]
		datagrams = append(datagrams, datagram)
	}

	return datagrams
}

func ParseTCPSegment(datagrams [][]byte) DecodedImage {
	image := DecodedImage{}
	for _, datagram := range datagrams {
		sourcePort := binary.BigEndian.Uint16(datagram[:2])
		destPort := binary.LittleEndian.Uint16(datagram[2:4])
		fmt.Println("------ TCP Segment -----")
		fmt.Printf("SOURCE PORT: %d DEST:%d \n", sourcePort, destPort)
		headerLength := (datagram[12] >> 4) * 4

    if sourcePort == 80 {

		sequenceNumber := binary.BigEndian.Uint32(datagram[4:8])
		httpData := datagram[headerLength:]
		imagePacket := DecodedPacket{order: int(sequenceNumber), data: httpData}
		image.packets = append(image.packets, imagePacket)
    }

	}

	return image

}

func CreateImage(imageData DecodedImage) {
	sort.Sort(imageData)
	for _, segment := range imageData.packets {
		fmt.Printf("Segment Number: %d - Length:%d\n", segment.order, len(segment.data))
	}
	image, err := os.Create("butts.jpg")
	if err != nil {
		fmt.Println("Error writing picture: ", err)
		return
	}
	defer image.Close()

	combinedData := make([][]byte, len(imageData.packets))

	for _, packet := range imageData.packets {
		combinedData = append(combinedData, packet.data)
	}

  response := bytes.Join(combinedData, []byte{})

	bodyData := ExtractHttpData(response)
	_, err = image.Write(bodyData)
	if err != nil {
		fmt.Println("Error writing body data:", err)
		image.Sync()
	}
}

func ExtractHttpData(data []byte) []byte {
	breakIndex := bytes.Index(data, []byte("\r\n\r\n"))
	fmt.Println("HI IM THE BREAK INDEX:", breakIndex)
	if breakIndex == -1 {
		fmt.Println("Oh I had to do the other one")
		breakIndex = bytes.Index(data, []byte("\n\n"))
	}

	if breakIndex != -1 {
		headers := data[:breakIndex]
		body := data[breakIndex+4:]
		fmt.Printf("HTTP HEADERS %s\n", string(headers))
		return body
	}
	fmt.Println("I looked at a packet and didnt find a body : /(")
	return make([]byte, 0)
}

func DecodeIPType(data []byte) string {
	var ipCode uint16 = binary.BigEndian.Uint16(data)
	if ipCode == 0x0800 {
		return "IPv4"
	}
	if ipCode == 0x0806 {
		return "ARP"
	}
	if ipCode == 0x86DD {
		return "IPv6"
	}
	return "oops"
}

func Parse6ByteNumber(data []byte) uint64 {
	var out uint64
	for i, x := range data {
		out |= uint64(x) << uint64(i*8)
	}
	return out
}

func Parse4BitNumber(data byte) int {
	num := data & 0x0F
	return int(num)
}

func First4Bits(data byte) int {
	num := (data >> 4) & 0x0F
	return int(num)
}

func (di DecodedImage) Len() int {
	return len(di.packets)
}
func (di DecodedImage) Less(i, j int) bool {
	return di.packets[i].order < di.packets[j].order
}
func (di DecodedImage) Swap(i, j int) {
	di.packets[i], di.packets[j] = di.packets[j], di.packets[i]
}
