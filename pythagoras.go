package main

import (
	"fmt"
	"math"
)	

func main() {
	var sisi_A, sisi_B, sisi_C float64
	

	fmt.Print("Input sisi A:")
	fmt.Scanf("%f\n", &sisi_A)

	fmt.Print("Input sisi B:")
	fmt.Scanf("%f\n", &sisi_B)

	sisi_C = math.Sqrt((sisi_A*sisi_A) + (sisi_B*sisi_B))

	fmt.Println("sisi adalah: ", sisi_C)
	
}