package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "E0529D18025800ABCA6996534CB22E4C00FB48E233BAEC947A8AA010CE1249DB51A02CC7DB67EF33D4002AE6ACDC40101CF0449AE4D9E4C071802D400F84BD21CAF3C8F2C35295EF3E0A600848F77893360066C200F476841040401C88908A19B001FD35CCF0B40012992AC81E3B980553659366736653A931018027C87332011E2771FFC3CEEC0630A80126007B0152E2005280186004101060C03C0200DA66006B8018200538012C01F3300660401433801A6007380132DD993100A4DC01AB0803B1FE2343500042E24C338B33F5852C3E002749803B0422EC782004221A41A8CE600EC2F8F11FD0037196CF19A67AA926892D2C643675A0C013C00CC0401F82F1BA168803510E3942E969C389C40193CFD27C32E005F271CE4B95906C151003A7BD229300362D1802727056C00556769101921F200AC74015960E97EC3F2D03C2430046C0119A3E9A3F95FD3AFE40132CEC52F4017995D9993A90060729EFCA52D3168021223F2236600ECC874E10CC1F9802F3A71C00964EC46E6580402291FE59E0FCF2B4EC31C9C7A6860094B2C4D2E880592F1AD7782992D204A82C954EA5A52E8030064D02A6C1E4EA852FE83D49CB4AE4020CD80272D3B4AA552D3B4AA5B356F77BF1630056C0119FF16C5192901CEDFB77A200E9E65EAC01693C0BCA76FEBE73487CC64DEC804659274A00CDC401F8B51CE3F8803B05217C2E40041A72E2516A663F119AC72250A00F44A98893C453005E57415A00BCD5F1DD66F3448D2600AC66F005246500C9194039C01986B317CDB10890C94BF68E6DF950C0802B09496E8A3600BCB15CA44425279539B089EB7774DDA33642012DA6B1E15B005C0010C8C917A2B880391160944D30074401D845172180803D1AA3045F00042630C5B866200CC2A9A5091C43BBD964D7F5D8914B46F040"

	packet, _ := parsePacket(toBinary(input))

	fmt.Println(sumVersions(packet))
	fmt.Println(packet.getValue())
}

func sumVersions(p packet) int {
	acc := 0
	q := []packet{p}
	for len(q) != 0 {
		e := q[0]
		q = q[1:]
		acc += e.getVersion()
		q = append(q, e.getChildren()...)
	}

	return acc
}

func toBinary(input string) string {
	res := ""
	for _, c := range input {
		i, _ := strconv.ParseInt(string(c), 16, 64)
		b := strconv.FormatInt(i, 2)
		res += appendLeadingZeros(b)
	}
	return res
}

func appendLeadingZeros(input string) string {
	res := input
	for true {
		if len(res) == 4 {
			return res
		}

		res = "0" + res
	}

	return input
}

func parsePacket(bits string) (p packet, tail string) {
	version, _ := strconv.ParseInt(bits[:3], 2, 64)
	id, _ := strconv.ParseInt(bits[3:6], 2, 64)

	switch id {
	case 4:
		value, rest := parseLiteral(bits[6:])
		p = literal{version: int(version), value: value}
		tail = rest
		return
	default:
		packets, rest := parseSubpackets(bits[6:])
		p = operator{version: int(version), operatorType: operatorId2Type(int(id)), subpackets: packets}
		tail = rest
		return
	}
}

func operatorId2Type(operatorId int) OperatorType {
	switch operatorId {
	case 0:
		return sum
	case 1:
		return product
	case 2:
		return min
	case 3:
		return max
	case 5:
		return gt
	case 6:
		return lt
	case 7:
		return et
	default:
		return unknown
	}
}

func parseSubpackets(bits string) (subpackets []packet, tail string) {
	lengthTypeId := string(bits[0])
	if lengthTypeId == "0" {
		subpacketBitLength, _ := strconv.ParseInt(bits[1:(1+15)], 2, 64)
		subpackets = parseSiblings(bits[(1 + 15):(1 + 15 + subpacketBitLength)])
		tail = bits[(1 + 15 + subpacketBitLength):]
		return
	} else {
		numberOfSupbackets, _ := strconv.ParseInt(bits[1:(1+11)], 2, 64)
		subpackets, tail = parseNSiblings(bits[(1+11):], int(numberOfSupbackets))
		return
	}
}

func parseSiblings(bits string) []packet {
	ps := make([]packet, 0)
	rest := bits
	for true {
		p, r := parsePacket(rest)
		rest = r
		ps = append(ps, p)
		if rest == "" {
			break
		}
	}

	return ps
}

func parseNSiblings(bits string, n int) (ps []packet, tail string) {
	res := make([]packet, 0)
	rest := bits
	for i := 0; i < n; i++ {
		p, r := parsePacket(rest)
		rest = r
		res = append(res, p)
	}
	return res, rest
}

func parseLiteral(bits string) (value int, tail string) {
	res := ""
	for i := 0; true; i++ {
		if string(bits[i*5]) == "1" {
			res += bits[(i*5 + 1):(i*5 + 5)]
		} else {
			res += bits[(i*5 + 1):(i*5 + 5)]

			v, _ := strconv.ParseInt(res, 2, 64)
			value = int(v)
			tail = bits[(i*5 + 5):]
			return
		}
	}

	return -1, ""
}

type packet interface {
	getVersion() int
	getValue() int
	getChildren() []packet
}

type literal struct {
	version int
	value   int
}

func (l literal) getVersion() int {
	return l.version
}

func (l literal) getValue() int {
	return l.value
}

func (l literal) getChildren() []packet {
	return make([]packet, 0)
}

type operator struct {
	version      int
	operatorType OperatorType
	subpackets   []packet
}

type OperatorType int

const (
	unknown OperatorType = iota
	sum     OperatorType = iota
	product OperatorType = iota
	min     OperatorType = iota
	max     OperatorType = iota
	gt      OperatorType = iota
	lt      OperatorType = iota
	et      OperatorType = iota
)

func (o operator) getVersion() int {
	return o.version
}

func (o operator) getValue() int {
	switch o.operatorType {
	case sum:
		acc := 0
		for _, ch := range o.getChildren() {
			acc += ch.getValue()
		}
		return acc
	case product:
		acc := 1
		for _, ch := range o.getChildren() {
			acc *= ch.getValue()
		}
		return acc
	case min:
		min := o.getChildren()[0].getValue()
		for _, ch := range o.getChildren() {
			v := ch.getValue()
			if v < min {
				min = v
			}
		}
		return min
	case max:
		max := o.getChildren()[0].getValue()
		for _, ch := range o.getChildren() {
			v := ch.getValue()
			if v > max {
				max = v
			}
		}
		return max
	case gt:
		v1 := o.getChildren()[0].getValue()
		v2 := o.getChildren()[1].getValue()
		if v1 > v2 {
			return 1
		} else {
			return 0
		}
	case lt:
		v1 := o.getChildren()[0].getValue()
		v2 := o.getChildren()[1].getValue()
		if v1 < v2 {
			return 1
		} else {
			return 0
		}
	case et:
		v1 := o.getChildren()[0].getValue()
		v2 := o.getChildren()[1].getValue()
		if v1 == v2 {
			return 1
		} else {
			return 0
		}
	default:
		return -1
	}
}

func (o operator) getChildren() []packet {
	return o.subpackets
}
