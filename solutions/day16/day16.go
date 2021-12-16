package day16

import (
	"advent-calendar/utils"
	"errors"
	"fmt"
	"strconv"
)

type Data string

type Day struct {
	data Data
}

func getHexMap() map[string]string {
	return map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
}

type Header struct {
	Version int
	TypeId  int
}

type OperatorLength struct {
	LengthTypeId int
	Length       int
}

type Packet struct {
	Header         Header
	SubPackets     []Packet
	Literal        int
	OperatorLength OperatorLength
}

func loadData(fname string) Data {
	return Data(utils.RawFileString(fname))
}

func NewDay(fname string) *Day {
	data := loadData(fname)
	return &Day{
		data: data,
	}
}

func decodeBit(hex string) string {
	out := ""
	for _, char := range hex {
		out += getHexMap()[string(char)]
	}
	return out
}

func getVersion(binString string) (int, string) {
	versionString, left := binString[0:3], binString[3:]
	version, err := strconv.ParseInt(versionString, 2, 64)
	utils.Check(err)
	return int(version), left
}

func parseBinaryInt(bitString string) int {
	num, err := strconv.ParseInt(bitString, 2, 64)
	utils.Check(err)
	return int(num)
}

func getTypeId(binString string) (int, string) {
	typeString, left := binString[0:3], binString[3:]
	typeNum, err := strconv.ParseInt(typeString, 2, 64)
	utils.Check(err)
	return int(typeNum), left
}

func parseLiteral(binString string) (int, string) {
	i := 0
	literalString := ""
	done := false
	for {
		if done || i >= len(binString) {
			break
		}
		section := binString[i : i+5]
		if string(section[0]) == "0" {
			done = true
		}
		literalString += section[1:]
		i += 5
	}
	num, err := strconv.ParseInt(literalString, 2, 64)
	utils.Check(err)
	return int(num), binString[i:]
}

func getLengthTypeId(bitString string) (int, string) {
	lengthTypeString, out := string(bitString[0]), bitString[1:]
	lengthType, err := strconv.ParseInt(lengthTypeString, 2, 64)
	utils.Check(err)
	return int(lengthType), out
}

func getTotalLengthOfSubpackets(bitString string) (int, string) {
	totalLength, out := bitString[0:15], bitString[15:]
	length := parseBinaryInt(totalLength)
	return length, out
}

func getNumberOfSubpackets(bitString string) (int, string) {
	numString, out := bitString[0:11], bitString[11:]
	return parseBinaryInt(numString), out
}

func parseSubpacketByLength(bitString string, length int) ([]Packet, string) {
	initialLength := len(bitString)
	out := bitString
	var packets []Packet
	for {
		nextPacket, nextOut := parsePacket(out)
		packets = append(packets, nextPacket)
		out = nextOut
		if initialLength-len(out) >= length {
			return packets, out
		}
	}
}

func parseSubpacketByNumber(bitString string, numPackets int) ([]Packet, string) {
	var packets []Packet
	out := bitString
	for {
		nextPacket, nextOut := parsePacket(out)
		packets = append(packets, nextPacket)
		out = nextOut
		if len(packets) >= numPackets {
			return packets, out
		}
	}
}

func parseOperator(packet Packet, bitString string) (Packet, string) {
	out := bitString
	lengthType, out := getLengthTypeId(out)
	if lengthType == 0 {
		totalLength, out := getTotalLengthOfSubpackets(out)
		packet.OperatorLength = OperatorLength{
			LengthTypeId: lengthType,
			Length:       totalLength,
		}
		subPackets, out := parseSubpacketByLength(out, totalLength)
		packet.SubPackets = subPackets
		return packet, out
	} else {
		packetLength, out := getNumberOfSubpackets(out)
		packet.OperatorLength = OperatorLength{
			LengthTypeId: lengthType,
			Length:       packetLength,
		}
		subPackets, out := parseSubpacketByNumber(out, packetLength)
		packet.SubPackets = subPackets
		return packet, out
	}
}

func parsePacket(bitString string) (Packet, string) {
	packet := Packet{}
	out := bitString
	// get header
	version, out := getVersion(out)
	typeId, out := getTypeId(out)

	packet.Header = Header{
		Version: version,
		TypeId:  typeId,
	}

	if typeId == 4 {
		literal, nextOut := parseLiteral(out)
		packet.Literal = literal
		return packet, nextOut
	} else {
		packet, nextOut := parseOperator(packet, out)
		return packet, nextOut
	}
}

func printPacket(packet Packet) {
	fmt.Println("---header---")
	fmt.Printf("version: %d typeId: %d\n", packet.Header.Version, packet.Header.TypeId)
	if packet.Header.TypeId == 4 {
		fmt.Printf("literal: %d\n", packet.Literal)
	} else {
		fmt.Println("---subpackets---")
		for _, subpacket := range packet.SubPackets {
			printPacket(subpacket)
		}
	}
}

func countVersions(packet Packet) int {
	sum := 0
	sum += packet.Header.Version
	for _, subpacket := range packet.SubPackets {
		sum += countVersions(subpacket)
	}
	return sum
}

func executeOperations(packet Packet) int {
	switch packet.Header.TypeId {
	case 0:
		// sum
		result := 0
		for _, subpacket := range packet.SubPackets {
			result += executeOperations(subpacket)
		}
		return result
	case 1:
		// prod
		result := 1
		for _, subpacket := range packet.SubPackets {
			result *= executeOperations(subpacket)
		}
		return result
	case 2:
		// min
		result := -1
		for _, subpacket := range packet.SubPackets {
			val := executeOperations(subpacket)
			if result == -1 || val < result {
				result = val
			}
		}
		return result
	case 3:
		// max
		result := -1
		for _, subpacket := range packet.SubPackets {
			val := executeOperations(subpacket)
			if result == -1 || val > result {
				result = val
			}
		}
		return result
	case 4:
		// literal
		return packet.Literal
	case 5:
		// greater than
		if executeOperations(packet.SubPackets[0]) > executeOperations(packet.SubPackets[1]) {
			return 1
		} else {
			return 0
		}
	case 6:
		// less than
		if executeOperations(packet.SubPackets[0]) < executeOperations(packet.SubPackets[1]) {
			return 1
		} else {
			return 0
		}
	case 7:
		// equal to
		if executeOperations(packet.SubPackets[0]) == executeOperations(packet.SubPackets[1]) {
			return 1
		} else {
			return 0
		}
	default:
		panic(errors.New(fmt.Sprintf("unknown typeId: %d!", packet.Header.TypeId)))
	}
}

func (d *Day) RunPart1() {
	out := decodeBit(string(d.data))
	packet, _ := parsePacket(out)
	sum := countVersions(packet)
	fmt.Printf("Part 1: %d\n", sum)
}

func (d *Day) RunPart2() {
	out := decodeBit(string(d.data))
	packet, _ := parsePacket(out)
	result := executeOperations(packet)
	fmt.Printf("Part 2: %d\n", result)
}
