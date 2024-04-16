package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type PacketData struct {
	Datetime1 DateTime  `json:"datetime1"`
	Datetime2 time.Time `json:"datetime2"`
	Datetime3 time.Time `json:"datetime3"`
}

type DateTime struct {
	CurTime time.Time
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	ret := d.CurTime.In(jst).Format("2006-01-02T15:04:05Z07:00")
	return json.Marshal(ret)
}

func main() {
	packet := PacketData{
		Datetime1: DateTime{CurTime: time.Now()},
		Datetime2: time.Now(),
		Datetime3: time.Now().UTC(),
	}
	j, _ := json.Marshal(packet)
	fmt.Println(string(j))
}
