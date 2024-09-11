package main

import "fmt"

func main() {
	var nama string
	var umur int

	fmt.Print("Input nama anda: ")
	fmt.Scanf("%s\n", &nama)

	fmt.Print("Input umur anda: ")
	fmt.Scanf("%d\n", &umur)

	fmt.Println("Hello", nama)
	fmt.Println("Umur anda: ", umur )

}