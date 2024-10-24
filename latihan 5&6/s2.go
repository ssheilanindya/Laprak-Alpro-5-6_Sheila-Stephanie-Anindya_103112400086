package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan n kerucut: ")
	fmt.Scan(&n)

	var jariJari, tinggi, volume, volumeTotal float64

	for j := 0; j < n; j++ {
		fmt.Print("Masukkan jari-jari kerucut: ", j+1)
		fmt.Scan(&jariJari)

		fmt.Print("Masukkan tinggi kerucuit: ", j+1)
		fmt.Scan(&tinggi)

		volume = (1.0 / 3.0) * math.Pi * math.Pow(jariJari, 2) * tinggi
		volumeTotal += volume

		fmt.Printf("Volume kerucut ke-%d: %.2f\n", j+1, volume)
	}
	fmt.Print(volumeTotal)
}
