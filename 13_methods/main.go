package main

import (
	"fmt"
)

const uSixTeenBitMax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16 // min 0 max 65535 (Number)
	brakePedal    uint16
	steeringWheel int16 // -32k / + 32k
	topSpeedKmh   float64
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / uSixTeenBitMax)
}

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / uSixTeenBitMax / kmhMultiple)

}

// pointer receivers methods (no return values but yes parameter)

func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

func main() {
	aCar := car{
		gasPedal:      65000,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKmh:   225.0,
	}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	aCar.newTopSpeed(180)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}
