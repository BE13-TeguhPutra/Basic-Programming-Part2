package main

import "fmt"

func Palindrome(input string) bool {
	// your code here
	var a []string
	var b []string
	x := 0

	for i := 0; i < len(input); i++ {
		a = append(a, string(input[i]))
	}
	for i := len(input) - 1; i >= 0; i-- {
		b = append(b, string(input[i]))
	}
	for i := 0; i < len(input); i++ {
		if a[i] == b[i] {
			x++
		}

	}
	return x == len(input)
}

func main() {
	fmt.Println(Palindrome("civic"))       // true
	fmt.Println(Palindrome("katak"))       // true
	fmt.Println(Palindrome("kasur rusak")) // true
	fmt.Println(Palindrome("kupu-kupu"))   // false
	fmt.Println(Palindrome("lion"))        // false
}
