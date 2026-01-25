package parentpattern

import "time"

type Packet interface {
	Date() time.Time
	Type() string
}

type packet struct {
	dateField time.Time
	typeField string
}

func (p packet) Date() time.Time {
	return p.dateField
}

func (p packet) Type() string {
	return p.typeField
}

func NewPacket(date time.Time, typeStr string) packet {
	return packet{
		dateField: date,
		typeField: typeStr,
	}
}

type TypeAPacket struct {
	Packet
	DetailA string
}

type TypeBPacket struct {
	Packet
	DetailB int
}
