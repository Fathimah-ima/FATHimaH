package main

import "fmt"

func main () {
	var nama, alamat string = "ima", "Kudus Jawa Tengah"
	var berkacamata bool = false
	var umur int = 19
	var minuskiri, minuskanan float32 = 0.0, 0.0
	var minustotal float32 = minuskiri + minuskanan
	
	fmt.Println("Nama =", nama)
	fmt.Println("Umur =", umur)
	fmt.Println("alamat =", alamat)
	fmt.Println("berkacamata =", berkacamata)
	fmt.Println("minus kiri :", minuskiri)
	fmt.Println("minus kanan :", minuskanan)
	fmt.Println("total minus :", minustotal)

}