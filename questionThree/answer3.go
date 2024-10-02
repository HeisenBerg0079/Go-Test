package main

import (
	"fmt"
	"math"
	"strconv"
)

type Car struct {
    wheelCircumference float64 
    roundsPerLiter     int
}

func NewCar() *Car {
    return &Car{
        wheelCircumference: 95,
        roundsPerLiter:     13000,
    }
}

func (c *Car) CalculateDistance(gasLiters float64) string {
    totalRounds := float64(c.roundsPerLiter) * gasLiters
    distanceInKm := int(math.Floor(totalRounds * (c.wheelCircumference / 100000)))
		strVal := strconv.Itoa(distanceInKm)
    return strVal + " KM"
}

func main() {
    car := NewCar()
    gasLiters := 10.0
    distance := car.CalculateDistance(gasLiters)
    fmt.Printf("Distance: %s \n", distance)
}