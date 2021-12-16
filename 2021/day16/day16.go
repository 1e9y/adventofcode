package day16

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
)

var versionSum int

// 0 = 0000
// 1 = 0001
// 2 = 0010
// 3 = 0011
// 4 = 0100
// 5 = 0101
// 6 = 0110
// 7 = 0111
// 8 = 1000
// 9 = 1001
// A = 1010
// B = 1011
// C = 1100
// D = 1101
// E = 1110
// F = 1111

// 100 010 1 00000000001 001010100000000001101010000000000000101111010001111000
// 4   2   1           1
// v 4

// 001 010 1 00000000001 101010000000000000101111010001111000
//   1   2 1           1
// v 5

// 101 010 0 000000000001011 11010001111 000
//   5   2                11
// v 10

// 110 100 01111 000
//   6   4
// v16

// 011 000 1 00000000010 00000000000000000101100001000101010110001011001000100000000010000100011000111000110100
//  v3   0             2

// 1:
// 000 000 0 000000000010110 0001000101010110001011 001000100000000010000100011000111000110100
//  v0   0                22

// 000 100 01010 10110001011
//  v0   4

// 101 100 01011
//  v5   4

// 2:
// 001 000 1 00000000010 000100011000111000110100
//  v1   0             2

// 000 100 01100 0111000110100
//  v0

// 011 100 0110100
//  v3

func decode(input string) int {
	// fmt.Println("+ input")
	// fmt.Println(input)
	builder := strings.Builder{}
	for _, c := range input {
		res, err := strconv.ParseInt(string(c), 16, 8)
		if err != nil {
			panic(err)
		}
		builder.WriteString(fmt.Sprintf("%04b", res))
	}
	bits := builder.String()
	// fmt.Println("+ bits")
	// fmt.Println(bits)

	packet, _ := readPackage(bits)
	println(packet)
	// pack, rams := readPackage(bits)
	// fmt.Println("+ package")
	// fmt.Printf("%#v\n", pack)
	// fmt.Println("+ remains")
	// fmt.Printf("%#v\n", rams)
	return calculate(packet)
}

func calculate(packet *Package) int {
	switch packet.typeID {
	case 0:
		sum := 0
		for _, p := range packet.subpackets {
			sum += calculate(p)
		}
		return sum
	case 1:
		product := 1
		for _, p := range packet.subpackets {
			product *= calculate(p)
		}
		return product
	case 2:
		value := math.MaxInt64
		for _, p := range packet.subpackets {
			nextvalue := calculate(p)
			if nextvalue < value {
				value = nextvalue
			}
		}
		return value
	case 3:
		value := 0
		for _, p := range packet.subpackets {
			nextvalue := calculate(p)
			if nextvalue > value {
				value = nextvalue
			}
		}
		return value
	case 4:
		return int(packet.value)
	case 5:
		if calculate(packet.subpackets[0]) > calculate(packet.subpackets[1]) {
			return 1
		}
		return 0
	case 6:
		if calculate(packet.subpackets[0]) < calculate(packet.subpackets[1]) {
			return 1
		}
		return 0
	case 7:
		if calculate(packet.subpackets[0]) == calculate(packet.subpackets[1]) {
			return 1
		}
		return 0
	}
	panic("bad stuff")
}

type Package struct {
	version     int64
	typeID      int64
	lengthType  int64
	lengthValue int64
	bits        string
	subpackets  []*Package
	value       int64
}

func readPackage(bits string) (*Package, string) {
	// println("PARSING", bits)
	pack := Package{
		subpackets: make([]*Package, 0),
	}
	version, err := strconv.ParseInt(bits[:3], 2, 8)
	if err != nil {
		panic(err)
	}
	pack.version = version
	bits = bits[3:]
	versionSum += int(version)

	id, err := strconv.ParseInt(bits[:3], 2, 8)
	if err != nil {
		panic(err)
	}
	pack.typeID = id
	bits = bits[3:]

	// println("TYPE", id, "VERSION", version)

	if id == 4 {
		pack.lengthType = -1
		pack.bits = bits

		subbits := ""
		for {
			group := bits[:5]
			bits = bits[5:]
			subbits += group[1:]
			if group[0] == '0' {
				break
			}
		}
		value, err := strconv.ParseInt(subbits, 2, 64)
		if err != nil {
			panic(err)
		}
		pack.value = value
		return &pack, bits
	}

	lenghtType, err := strconv.ParseInt(bits[:1], 2, 8)
	if err != nil {
		panic(err)
	}
	pack.lengthType = lenghtType
	bits = bits[1:]

	if lenghtType == 0 {
		lenghtValue, err := strconv.ParseInt(bits[:15], 2, 64)
		if err != nil {
			panic(err)
		}
		// total length
		pack.lengthValue = lenghtValue
		bits = bits[15:]
		pack.bits = bits

		origLen := len(bits)
		consumed := 0
		var lastBits string = bits

		var subpacket *Package
		// for consumed <= int(lenghtValue) {
		for consumed < int(lenghtValue) {
			subpacket, lastBits = readPackage(bits)
			// fmt.Println(">", origLen, lenghtValue, lastBits, consumed)
			consumed += (origLen - len(lastBits))
			origLen = len(lastBits)
			pack.subpackets = append(pack.subpackets, subpacket)
			// fmt.Printf("> %#v\n", subpacket)
			bits = lastBits
		}
		// println("after readling remained:", lastBits)
		// fmt.Printf("> %#v\n", subpacket)
		return &pack, lastBits
	}
	if lenghtType == 1 {
		lenghtValue, err := strconv.ParseInt(bits[:11], 2, 64)
		if err != nil {
			panic(err)
		}

		// count
		pack.lengthValue = lenghtValue
		pack.bits = bits
		bits = bits[11:]

		var lastBits string
		var subpacket *Package

		i := lenghtValue
		for i > 0 {
			// println()
			// println(i, bits)
			subpacket, lastBits = readPackage(bits)
			pack.subpackets = append(pack.subpackets, subpacket)
			bits = lastBits
			// fmt.Printf("> %v %#v\n", i, subpacket)
			i--
		}

		return &pack, lastBits
	}

	panic("bad input")
}

func A(input *challenge.Challenge) int {
	versionSum = 0
	decode(<-input.Lines())
	return versionSum
}

func B(input *challenge.Challenge) int {
	versionSum = 0
	return decode(<-input.Lines())
}
