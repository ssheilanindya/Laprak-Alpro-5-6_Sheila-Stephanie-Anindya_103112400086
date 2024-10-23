package main

import "fmt"

func main() {
	var basis, pangkat int

	fmt.Print("Masukkan bilangan pertama (basis): ")
	fmt.Scan(&basis)

	fmt.Print("Masukkan biangan kedua (pangkat): ")
	fmt.Scan(&pangkat)

	hasil := 1
	for j := 0; j < pangkat; j++ {
		hasil *= basis
	}
	fmt.Print(basis, pangkat, hasil)
}
