package main

import "fmt"

var c, python, java bool

func main() {
	oneString := "texto1"
	anotherString := "texto2"

	fmt.Println("Textos originais: ", oneString, anotherString)

	oneString, anotherString = swapStrings(oneString, anotherString)

	fmt.Println("Textos trocados: ", oneString, anotherString)

	var fahrenheitTemp, kelvinTemp float64

	fahrenheitTemp, kelvinTemp = getFahrenheitAndKelvinByCelsius(13)

	fmt.Println("Temperatura (13 C°) para Fahrenheit e Kelvin:", fahrenheitTemp, kelvinTemp)

	fmt.Println("Essa linguagem é c, python ou java? ", c, python, java)
}

func swapStrings(firstString, secondString string) (string, string) {
	return secondString, firstString
}

func getFahrenheitAndKelvinByCelsius(celsiusTemp float64) (fahrenheitTemp float64, kelvinTemp float64) {
	fahrenheitTemp = (celsiusTemp * 9 / 5) + 32
	kelvinTemp = celsiusTemp + 273.15
	return
}
