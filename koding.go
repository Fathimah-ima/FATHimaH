package main

import (
	"fmt"
	"os"
)

// deklarasi variabel
var (
	Username = "Fathimah"
	Password = "2406423805"
	Menu     int
	Histori  []string
)

type Pengguna struct {
	Username   string
	NPM   string
	Gender string
	Makanan string
	Minuman string
}

type Buku struct {
	Nama   string
	Jumlah int
}

var daftarBuku = []Buku{ // <-- struct
	{Nama: "Pemrograman", Jumlah: 10}, //array 0 buku
	{Nama: "Film", Jumlah: 5}, //array 1 buku
	{Nama: "Printing", Jumlah: 20},  //array 2 buku
}

func LihatInformasiPenggunaProgram() {
	data := Pengguna{
		Username:   "Fathimah",
		NPM:   "2406423805",
		Gender: "Perempuan",
		Makanan: "Mie",
		Minuman: "Es Jeruk",
	}
	fmt.Println("============================================================================")
	fmt.Println("Informasi Pengguna Program")
	fmt.Println("Username: ", data.Username)
	fmt.Println("NPM: ", data.NPM)
	fmt.Println("Jenis Kelamin: ", data.Gender)
	fmt.Println("Makanan Favorit: ", data.Makanan)
	fmt.Println("Minuman Favorit: ", data.Minuman)
	fmt.Println("============================================================================")
}

func LihatDaftarBuku() {
	fmt.Println("============================================================================")
	fmt.Println("Daftar Buku")
	for i, buku := range daftarBuku {
		fmt.Printf("%v. Nama Buku: %s \n   Jumlah: %v \n",   i+1, buku.Nama, buku.Jumlah)
	}
	fmt.Println("============================================================================")
}

func TambahDaftarBuku() {
	// var newBuku Buku = Buku{}
	// var newBuku Buku
	newBuku := Buku{}
	fmt.Println("============================================================================")
	fmt.Println("Tambah Daftar Buku")
	fmt.Print("Buku: ")
	fmt.Scanln(&newBuku.Nama)
	fmt.Print("Jumlah: ")
	fmt.Scanln(&newBuku.Jumlah)
	fmt.Println("============================================================================")

	daftarBuku = append(daftarBuku, newBuku)
	Histori = append(Histori, fmt.Sprintf("Menambahkan buku %s sejumlah %v", newBuku.Nama, newBuku.Jumlah))
	fmt.Println("Buku berhasil ditambahkan")
	fmt.Println("============================================================================")
}

func TambahPeminjamanBuku() {
	fmt.Println("Pinjam Buku")
	fmt.Println("1. Pemrograman (10)\n2. Film (5)\n3. Printing (20)")
	var pilihBuku, jmlBuku int
	fmt.Print("Masukkan Nomor Pinjaman Buku (1,2,3): ")
	fmt.Scanln(&pilihBuku)
	fmt.Print("Jumlah Pinjaman: ")
	fmt.Scanln(&jmlBuku)

	for i := 1; i < len(daftarBuku); i++ {
		if i == pilihBuku {
			daftarBuku[i-1].Jumlah -= jmlBuku
			Histori = append(Histori, fmt.Sprintf("Meminjam buku %s sejumlah %v", daftarBuku[i-1].Nama, jmlBuku))
		}
	}
	fmt.Println("============================================================================")
	fmt.Println("=Selamat Peminjaman Anda Berhasil=")
	fmt.Println("============================================================================")
}

func HistoriPeminjamanBuku() {
	fmt.Println("============================================================================")
	fmt.Println("Histori Peminjaman Buku")
	if len(Histori) == 0 {
		fmt.Println("Data histori tidak ditemukan")
	} else {
		for _, histori := range Histori {
			fmt.Println(histori)
		}
	}
	fmt.Println("============================================================================")
}

func main() {
	//variabel input
	var inputUsername, inputPassword string

	//pesan pembuka
	fmt.Println("============================================================================")
	fmt.Println("=Selamat Datang di Perpustakaan Vokasi=")
	fmt.Println("============================================================================")

	//input username
	fmt.Print("Silahkan Input Username: ")
	fmt.Scanf("%s\n", &inputUsername)

	//input password
	fmt.Print("Silahkan Input Password: ")
	fmt.Scanf("%s\n", &inputPassword)
	fmt.Println("============================================================================")

	//validasi username &password jika salah keduanya program terhenti
	if inputUsername != Username || inputPassword != Password {
		fmt.Println("Password atau Username Anda, SALAH!")
		return
	}

	for {

		fmt.Println("Login Sukses!!!!")
		fmt.Println("Menu Program: ")
		fmt.Println("1. Lihat Informasi Pengguna Program")
		fmt.Println("2. Lihat Daftar Buku")
		fmt.Println("3. Tambah Daftar Buku")
		fmt.Println("4. Tambah Peminjaman Buku")
		fmt.Println("5. Histori Peminjaman Buku")
		fmt.Println("6. Keluar dari Program")

		//input pilihan menu
		var choice int
		fmt.Print("Input Menu yang anda inginkan: ")
		fmt.Scanf("%v\n", &choice)
		fmt.Println("============================================================================")

		switch choice {
		case 1:
			// fmt.Println("Lihat Informasi Pengguna Program")
			LihatInformasiPenggunaProgram()
		case 2:
			// fmt.Println("Lihat daftar boneka")
			LihatDaftarBuku()
		case 3:
			// fmt.Println("Tambah daftar boneka")
			TambahDaftarBuku()
		case 4:
			// fmt.Println("Tambah peminjaman boneka")
			TambahPeminjamanBuku()
		case 5:
			// fmt.Println("Histori peminjaman boneka")
			HistoriPeminjamanBuku()
		case 6:
			// fmt.Println("Keluar dari program")
			fmt.Println("=Terima kasih telah Mengunjungi Perpustakaan Vokasi=")
			fmt.Println("============================================================================")
			os. Exit(0)
		default:
			fmt.Println("Pilihan anda tidak valid")
		}
	}

}