package main

import (
	"fmt"
)

func PrimeNumber(number int) bool {
	// your code here

	var y int = 0
	for i := 2; i < number; i++ {
		if number%i == 0 {
			y++
		}
	}
	return (y == 0)

}

func main() {
	fmt.Println(PrimeNumber(2))
	fmt.Println(PrimeNumber(7))
	fmt.Println(PrimeNumber(10))
}
