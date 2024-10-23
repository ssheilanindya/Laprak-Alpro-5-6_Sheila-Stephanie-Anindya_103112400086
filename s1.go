package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif :")
	fmt.Scan(&n)

	sum := 0
	for j := 1; j <= n; j++ {
		sum += j
	}
	fmt.Print(sum)
}
