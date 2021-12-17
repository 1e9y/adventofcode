package day16

import (
	"fmt"
	"math"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type Type int

const (
	TypeSum Type = iota
	TypeProduct
	TypeMinimum
	TypeMaximum
	TypeLiteral
	TypeGreaterThan
	TypeLessThan
	TypeEqualTo
)

type Packet struct {
	version    int
	typeID     Type
	value      int
	lengthType int
	length     int
	subpackets []*Packet
}

func newPacketFromInput(input string) *Packet {
	builder := strings.Builder{}
	for _, c := range input {
		res := util.MustParseInt(string(c), 16)
		builder.WriteString(fmt.Sprintf("%04b", res))
	}
	bits := builder.String()

	packet, _ := readPacket(bits)
	return packet
}

func readPacket(bits string) (*Packet, string) {
	packet := Packet{
		lengthType: -1,
		subpackets: make([]*Packet, 0),
	}

	version := util.MustParseInt(bits[:3], 2)
	packet.version = version
	bits = bits[3:]

	id := util.MustParseInt(bits[:3], 2)
	packet.typeID = Type(id)
	bits = bits[3:]

	if id == int(TypeLiteral) {
		tail := packet.readLiteral(bits)
		return &packet, tail
	}

	lenghtType := util.MustParseInt(bits[:1], 2)
	packet.lengthType = lenghtType
	bits = bits[1:]

	switch lenghtType {
	case 0:
		tail := packet.readPacketByLength(bits)
		return &packet, tail
	case 1:
		tail := packet.readPacketByNumber(bits)
		return &packet, tail
	}

	panic(fmt.Errorf("bad packet: %s", bits))
}

func (packet *Packet) readLiteral(bits string) string {
	var group, value string
	for {
		group = bits[:5]
		bits = bits[5:]
		value += group[1:]
		if group[0] == '0' {
			break
		}
	}
	packet.value = util.MustParseInt(value, 2)
	return bits
}

func (packet *Packet) readPacketByLength(bits string) (tail string) {
	length := util.MustParseInt(bits[:15], 2)
	packet.length = length
	bits = bits[15:]

	var subpacket *Packet

	tail = bits
	rest := len(bits)
	consumed := 0
	for consumed < length {
		subpacket, tail = readPacket(bits)
		packet.subpackets = append(packet.subpackets, subpacket)
		consumed += (rest - len(tail))
		rest = len(tail)
		bits = tail
	}

	return tail
}

func (packet *Packet) readPacketByNumber(bits string) (tail string) {
	length := util.MustParseInt(bits[:11], 2)
	packet.length = length
	bits = bits[11:]

	var subpacket *Packet

	count := length
	for count > 0 {
		subpacket, tail = readPacket(bits)
		packet.subpackets = append(packet.subpackets, subpacket)
		bits = tail
		count--
	}

	return tail
}

func (packet *Packet) calculate() int {
	switch packet.typeID {
	case TypeSum:
		sum := 0
		for _, p := range packet.subpackets {
			sum += p.calculate()
		}
		return sum
	case TypeProduct:
		product := 1
		for _, p := range packet.subpackets {
			product *= p.calculate()
		}
		return product
	case TypeMinimum:
		value := math.MaxInt64
		for _, p := range packet.subpackets {
			nextvalue := p.calculate()
			if nextvalue < value {
				value = nextvalue
			}
		}
		return value
	case TypeMaximum:
		value := 0
		for _, p := range packet.subpackets {
			nextvalue := p.calculate()
			if nextvalue > value {
				value = nextvalue
			}
		}
		return value
	case TypeLiteral:
		return int(packet.value)
	case TypeGreaterThan:
		if packet.subpackets[0].calculate() > packet.subpackets[1].calculate() {
			return 1
		}
		return 0
	case TypeLessThan:
		if packet.subpackets[0].calculate() < packet.subpackets[1].calculate() {
			return 1
		}
		return 0
	case TypeEqualTo:
		if packet.subpackets[0].calculate() == packet.subpackets[1].calculate() {
			return 1
		}
		return 0
	}

	panic(fmt.Errorf("bad packet type: %v", packet.typeID))
}

func (packet *Packet) versions() (sum int) {
	sum += int(packet.version)
	for _, p := range packet.subpackets {
		sum += p.versions()
	}
	return
}

func A(input *challenge.Challenge) int {
	return newPacketFromInput(<-input.Lines()).versions()
}

func B(input *challenge.Challenge) int {
	return newPacketFromInput(<-input.Lines()).calculate()
}
