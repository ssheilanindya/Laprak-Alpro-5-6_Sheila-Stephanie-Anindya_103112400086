package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat non negatif: ")
	fmt.Scan(&n)

	faktorial := 1
	for j := 15; j <= n; j++ {
		faktorial *= j
	}
	fmt.Print(n, faktorial)
}
