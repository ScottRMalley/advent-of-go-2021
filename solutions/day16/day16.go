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
	bitString := ""
	for _, char := range hex {
		bitString += getHexMap()[string(char)]
	}
	return bitString
}

func parseBinaryInt(bitString string) int {
	num, err := strconv.ParseInt(bitString, 2, 64)
	utils.Check(err)
	return int(num)
}

func getVersion(bitString string) (int, string) {
	versionString, left := bitString[0:3], bitString[3:]
	return parseBinaryInt(versionString), left
}

func getTypeId(bitString string) (int, string) {
	typeString, left := bitString[0:3], bitString[3:]
	return parseBinaryInt(typeString), left
}

func parseLiteral(bitString string) (int, string) {
	i := 0
	literalString := ""
	done := false
	for {
		if done || i >= len(bitString) {
			break
		}
		section := bitString[i : i+5]
		if string(section[0]) == "0" {
			done = true
		}
		literalString += section[1:]
		i += 5
	}
	return parseBinaryInt(literalString), bitString[i:]
}

func getLengthTypeId(bitString string) (int, string) {
	return parseBinaryInt(string(bitString[0])), bitString[1:]
}

func getTotalLengthOfSubpackets(bitString string) (int, string) {
	return parseBinaryInt(bitString[0:15]), bitString[15:]
}

func getNumberOfSubpackets(bitString string) (int, string) {
	return parseBinaryInt(bitString[0:11]), bitString[11:]
}

func parseSubpacketByLength(bitString string, length int) ([]Packet, string) {
	initialLength := len(bitString)
	remainder := bitString
	var packets []Packet
	var packet Packet
	for {
		packet, remainder = parsePacket(remainder)
		packets = append(packets, packet)
		if initialLength-len(remainder) >= length {
			return packets, remainder
		}
	}
}

func parseSubpacketByNumber(bitString string, numPackets int) ([]Packet, string) {
	var packets []Packet
	var packet Packet
	remainder := bitString
	for {
		packet, remainder = parsePacket(remainder)
		packets = append(packets, packet)
		if len(packets) >= numPackets {
			return packets, remainder
		}
	}
}

func parseOperator(packet Packet, bitString string) (Packet, string) {
	remainder := bitString
	lengthType, remainder := getLengthTypeId(remainder)
	if lengthType == 0 {
		totalLength, remainder := getTotalLengthOfSubpackets(remainder)
		packet.OperatorLength = OperatorLength{
			LengthTypeId: lengthType,
			Length:       totalLength,
		}
		subPackets, remainder := parseSubpacketByLength(remainder, totalLength)
		packet.SubPackets = subPackets
		return packet, remainder
	} else {
		packetLength, remainder := getNumberOfSubpackets(remainder)
		packet.OperatorLength = OperatorLength{
			LengthTypeId: lengthType,
			Length:       packetLength,
		}
		subPackets, remainder := parseSubpacketByNumber(remainder, packetLength)
		packet.SubPackets = subPackets
		return packet, remainder
	}
}

func parsePacket(bitString string) (Packet, string) {
	packet := Packet{}
	remainder := bitString
	// get header
	version, remainder := getVersion(remainder)
	typeId, remainder := getTypeId(remainder)

	packet.Header = Header{
		Version: version,
		TypeId:  typeId,
	}

	if typeId == 4 {
		literal, out := parseLiteral(remainder)
		packet.Literal = literal
		return packet, out
	} else {
		packet, remainder = parseOperator(packet, remainder)
		return packet, remainder
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
	bitString := decodeBit(string(d.data))
	packet, _ := parsePacket(bitString)
	sum := countVersions(packet)
	fmt.Printf("Part 1: %d\n", sum)
}

func (d *Day) RunPart2() {
	bitString := decodeBit(string(d.data))
	packet, _ := parsePacket(bitString)
	result := executeOperations(packet)
	fmt.Printf("Part 2: %d\n", result)
}
