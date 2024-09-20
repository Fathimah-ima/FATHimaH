package main

import (
	"fmt"
)

func main() {

	// deklarasi variabel
	var nama_customer, nama_barang string
	var harga float32 = 25000
	var quantity int32
	var discount, total_harga float32

	for {

		// input nama customer
		fmt.Print("input nama customer: ")
		fmt.Scanf("%s\n", &nama_customer)

		// input nama barang
		fmt.Print("input nama barang: ")
		fmt.Scanf("%s\n", &nama_barang)

		// input quantity barang
		fmt.Print("input quantity barang: ")
		fmt.Scanf("%d\n", &quantity)

		// kondisi discount, kalau lebih dari 5 dapat 10%, selain itu 2%
		switch quantity {
		case 5 - 1000000000:
			discount = 0.1
		default:
			discount = 0.02
		}

		// proses
		sub_total := float32(quantity) * harga
		total_harga = sub_total - (discount * sub_total)

		// tampilkan hasil harga
		fmt.Println("=======================")
		fmt.Println("Total barang: ", quantity)
		fmt.Println("Harga: ", harga)
		fmt.Println("Subtotal: ", sub_total)
		fmt.Println("Total Diskon:", discount*sub_total)
		fmt.Println("Total: ", total_harga)

		var ulang string
		fmt.Println("=======================")
		fmt.Print("Apakah Anda ingin menghitung ulang? (y/n): ")
		fmt.Scanf("%s\n", &ulang)
		fmt.Println("=======================")

		if ulang != "y" && ulang != "Y" {
			break // Keluar dari loop jika tidak ingin menghitung ulang
		}

	}
}
