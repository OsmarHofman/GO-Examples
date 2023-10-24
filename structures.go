package main

import "fmt"

func main() {
	defer fmt.Println("Somas feitas")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("Soma usando for: ", sum)

	isSumEven(sum)

	sum = 2
	for sum < 50 {
		if rest := sum % 2; rest == 0 {
			sum += 1
		} else {
			sum += rest
		}
	}

	fmt.Println("Soma usando 'while': ", sum)

	isSumEven(sum)
}

func isSumEven(sum int) {
	switch isEven := sum % 2; isEven {
	case 0:
		fmt.Println("A soma é par.")
	case 1:
		fmt.Println("A soma é impar.")
	}
}
