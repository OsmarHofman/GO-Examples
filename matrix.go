package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var slice []int = primes[1:]
	fmt.Println(slice)

	slice = append(slice, 17)
	fmt.Println(slice)

	madeSlice := make([]int, 6)
	fmt.Println(madeSlice)

	sumAndPrimesCount := []int{0, 0}
	for _, v := range slice {
		sumAndPrimesCount[0] += v
	}
	sumAndPrimesCount[1] = len(primes)
	fmt.Println(sumAndPrimesCount)

}
