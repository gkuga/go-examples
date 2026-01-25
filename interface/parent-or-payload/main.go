package main

import (
	"fmt"
	"time"

	"github.com/gkuga/go-examples/interface/parent-or-payload/pkg/parentpattern"
	"github.com/gkuga/go-examples/interface/parent-or-payload/pkg/payloadpattern"
)

func main() {
	// Using Parent Pattern
	handleParentPatternPacket(&parentpattern.TypeAPacket{
		Packet: parentpattern.NewPacket(
			time.Now(),
			"TypeA",
		),
		DetailA: "Example Detail A",
	})

	handleParentPatternPacket(&parentpattern.TypeBPacket{
		Packet: parentpattern.NewPacket(
			time.Now(),
			"TypeB",
		),
		DetailB: 42,
	})

	// Using Payload Pattern
	handlePayloadPatternPacket(payloadpattern.Packet{
		Date: time.Now(),
		Type: "TypeA",
		Payload: &payloadpattern.TypeAPayload{
			DetailA: "Example Detail A",
		},
	})

	handlePayloadPatternPacket(payloadpattern.Packet{
		Date: time.Now(),
		Type: "TypeB",
		Payload: &payloadpattern.TypeBPayload{
			DetailB: 42,
		},
	})
}

func handleParentPatternPacket(p parentpattern.Packet) {
	fmt.Printf("Packet Date: %s, Type: %s\n", p.Date(), p.Type())
	switch pkt := p.(type) {
	case *parentpattern.TypeAPacket:
		fmt.Printf("Handling Type A Packet with DetailA: %s\n", pkt.DetailA)
	case *parentpattern.TypeBPacket:
		fmt.Printf("Handling Type B Packet with DetailB: %d\n", pkt.DetailB)
	default:
		fmt.Println("Unknown packet type")
	}
}

func handlePayloadPatternPacket(p payloadpattern.Packet) {
	fmt.Printf("Packet Date: %s, Type: %s\n", p.Date, p.Type)
	switch payload := p.Payload.(type) {
	case *payloadpattern.TypeAPayload:
		fmt.Printf("Handling Type A Payload with DetailA: %s\n", payload.DetailA)
	case *payloadpattern.TypeBPayload:
		fmt.Printf("Handling Type B Payload with DetailB: %d\n", payload.DetailB)
	default:
		fmt.Println("Unknown payload type")
	}
}
