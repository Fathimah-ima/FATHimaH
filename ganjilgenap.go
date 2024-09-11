package main

import "fmt"

func main() {
	var angka int16
	var informasiGanjilGenap string
	var habisTerbagiDua bool

	fmt.Print("Input Angka: ")
	fmt.Scanf("%d", &angka)

	habisTerbagiDua = (angka%2 == 0)

	if habisTerbagiDua == true {
		informasiGanjilGenap = "Genap"
	} else {
		informasiGanjilGenap = "Ganjil"
	}

	fmt.Println(informasiGanjilGenap)

}
