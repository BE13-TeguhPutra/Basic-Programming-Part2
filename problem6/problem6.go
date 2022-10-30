package main

import (
	"fmt"
	"strconv"
)

func FullPrima(n int) bool {
	var y int = 0
	var str_n = strconv.Itoa(n) //convert n menjadi string
	for i := 2; i < n; i++ {
		if n%i == 0 {
			y++
		}
	}

	if y == 0 {
		for i := 0; i <= len(str_n); i++ {
			konv1, err := strconv.Atoi(string(str_n[i]))
			for i := 2; i < konv1; i++ {
				if konv1%i == 0 {
					return false
				}

		}

	}
	return true

}

//write your code

func main() {
	fmt.Println(FullPrima(2))  // true
	fmt.Println(FullPrima(3))  // true
	fmt.Println(FullPrima(11)) // false
	fmt.Println(FullPrima(13)) // false
	fmt.Println(FullPrima(23)) // true
	fmt.Println(FullPrima(29)) // false
	fmt.Println(FullPrima(37)) // true
	fmt.Println(FullPrima(41)) // false
	fmt.Println(FullPrima(43)) // false
	fmt.Println(FullPrima(53)) // true

}
