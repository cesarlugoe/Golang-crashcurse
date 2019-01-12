package main

import (
	"fmt"
)

type car struct {
	gasPedal      uint16 // min 0 max 65535 (Number)
	brakePedal    uint16
	steeringWheel int16 // -32k / + 32k
	topSpeedKmh   float64
}

func main() {
	aCar := car{
		gasPedal:      22341,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKmh:   225.0,
	}

	fmt.Println(aCar.gasPedal)
}
