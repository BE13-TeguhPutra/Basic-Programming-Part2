package main

import "fmt"

func FaktorBilanganDesc(n int) string {
	// your code here
	var faktor string
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			faktor += fmt.Sprintln(i)

		}
	}
	return faktor
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilanganDesc(number))
}
