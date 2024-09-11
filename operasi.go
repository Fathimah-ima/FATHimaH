package main

import (
	"fmt"
)

func main() {
	var bilangan_1 int
	var bilangan_2 int
	var input_user int
	var hasil int

	fmt.Print("1. Penjumlahan\n2. Pengurangan\n3. Perkalian\n4. Pembagian\n")

	fmt.Print("Pilih Operasi (Angka):")
	fmt.Scanf("%d\n", &input_user)

	fmt.Print("Input bilangan 1: ")
	fmt.Scanf("%d\n", &bilangan_1)

	fmt.Print("Input bilangan 2: ")
	fmt.Scanf("%d\n", &bilangan_2)

	if input_user == 1 {
		hasil = bilangan_1 + bilangan_2
		fmt.Println("=====Penjumlahan=====")
		fmt.Print(bilangan_1, " + ", bilangan_2, " = ", hasil)
	} else if input_user == 2 {
		hasil = bilangan_1 - bilangan_2
		fmt.Println("=====Pengurangan=====")
		fmt.Print(bilangan_1, " - ", bilangan_2, " = ", hasil)
	} else if input_user == 3 {
		hasil = bilangan_1 * bilangan_2
		fmt.Println("=====Perkalian=====")
		fmt.Print(bilangan_1, " x ", bilangan_2, " = ", hasil)
	} else if input_user == 4 {
		hasil = bilangan_1 / bilangan_2
		fmt.Println("=====Pembagian=====")
		fmt.Print(bilangan_1, " : ", bilangan_2, " = ", hasil)
	}
}
