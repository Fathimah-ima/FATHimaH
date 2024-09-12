package main

import "fmt"

func main() {
	// deklarasi variabel
	var bilangan1 int
	var bilangan2 int
	var inputuser int
	var hasil int

	// jenis operasi hitung
	fmt.Print("1. penjumlahan\n2. pengurangan\n3. perkalian\n4. pembagian\n")
	fmt.Print("pilih jenis operasi hitung (angka 1-4): ")
	fmt.Scanf("%d\n", &inputuser)


	fmt.Print("input bilangan 1: ")
	fmt.Scanf("%d\n", &bilangan1)

	fmt.Print("input bilangan 2: ")
	fmt.Scanf("%d\n", &bilangan2)

	if inputuser == 1 {

		hasil = bilangan1 + bilangan2
		fmt.Println("===penjumlahan===")
		fmt.Print(bilangan1, " + ", bilangan2, " = ", hasil)

	} else if inputuser == 2 {

		hasil = bilangan1 - bilangan2
		fmt.Println("===pengurangan===")
		fmt.Print(bilangan1, " - ", bilangan2, " = ", hasil)

	} else if inputuser == 3 {

		hasil = bilangan1 * bilangan2
		fmt.Println("===perkalian===")
		fmt.Print(bilangan1, " x ", bilangan2, " = ", hasil)

	} else if inputuser == 4 {

		hasil = bilangan1 / bilangan2
		fmt.Print("===pembagian===\n")
		fmt.Println(bilangan1, " : ", bilangan2, " = ", hasil)

	} else {
		fmt.Print("kalkulator tidak mengerti input user")
	}
	
}