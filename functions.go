package main

import (
	"fmt"
	"math/cmplx"
)

var c, python, java, goLang = false, false, false, true

func main() {
	const CelsiusTemp = 13

	var (
		oneString                             = "texto1"
		anotherString                         = "texto2"
		squareRoot                 complex128 = cmplx.Sqrt(-5 + 12i)
		fahrenheitTemp, kelvinTemp float64
	)

	fmt.Println("Textos originais: ", oneString, anotherString)

	oneString, anotherString = swapStrings(oneString, anotherString)

	fmt.Println("Textos trocados: ", oneString, anotherString)

	fmt.Println("Tipo complexo (raiz quadrada):", squareRoot)

	fahrenheitTemp, kelvinTemp = getFahrenheitAndKelvinByCelsius(CelsiusTemp)

	formatTemp := fmt.Sprintf("Temperatura (%d C°) para Fahrenheit e Kelvin: %.2f %.2f", CelsiusTemp, fahrenheitTemp, kelvinTemp)

	fmt.Println(formatTemp)

	fmt.Println("Essa linguagem é c, python, java ou go? ", c, python, java, goLang)
}

func swapStrings(firstString, secondString string) (string, string) {
	return secondString, firstString
}

func getFahrenheitAndKelvinByCelsius(celsiusTemp float64) (fahrenheitTemp float64, kelvinTemp float64) {
	fahrenheitTemp = (celsiusTemp * 9 / 5) + 32
	kelvinTemp = celsiusTemp + 273.15
	return
}
