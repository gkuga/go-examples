package payloadpattern

import "time"

type Packet struct {
	Date    time.Time
	Type    string
	Payload Payload
}

type Payload any

type TypeAPayload struct {
	Payload
	DetailA string
}

type TypeBPayload struct {
	Payload
	DetailB int
}

type GenericsPayload[T Payload] struct {
	Date    time.Time
	Type    string
	Payload T
}
