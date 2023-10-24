package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Pi: ", round(math.Pi, 5))
	fmt.Println("Número gerado aleatóriamente: ", rand.Intn(500))
}

func round(valueToRound float64, decimalPlaces int) float64 {
	var decimalPlaceValue = math.Pow(10, float64(decimalPlaces))
	return math.Round(valueToRound*decimalPlaceValue) / decimalPlaceValue
}
