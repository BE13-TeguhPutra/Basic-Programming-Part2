package main

import "fmt"

func KonversiNilai(nilai int) string {
	// your code here
	if nilai >= 80 && nilai <= 100 {
		return "Nilai A"
	} else if nilai >= 65 && nilai <= 79 {
		return "Nilai B+"
	} else if nilai >= 50 && nilai <= 64 {
		return "Nilai B"
	} else if nilai >= 35 && nilai <= 49 {
		return "Nilai C"
	} else {
		return "Nilai D"
	}
}

func main() {
	var nilai int = 80

	fmt.Println(KonversiNilai(nilai))
}
