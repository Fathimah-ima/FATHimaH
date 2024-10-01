package main

import (
	"fmt"
	"os"
)

type Mahasiswa struct {
	Username  string
	NPM       string
	Saldo     int
	Transaksi []string
}

func LihatAkun(mahasiswa Mahasiswa) {
	fmt.Println("=== Informasi Akun ===")
	fmt.Printf("Nama: %s\n", mahasiswa.Username)
	fmt.Printf("NPM: %s\n", mahasiswa.NPM)
}

func LihatSaldo(mahasiswa Mahasiswa) {
	fmt.Println("=== Informasi Saldo ===")
	fmt.Printf("Saldo saat ini: Rp%d\n", mahasiswa.Saldo)
}

func TambahSaldo(mahasiswa *Mahasiswa) {
	var jumlah int
	fmt.Println("=== Tambah Saldo ===")
	fmt.Print("Masukkan jumlah saldo yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)

	if jumlah > 0 {
		mahasiswa.Saldo += jumlah
		mahasiswa.Transaksi = append(mahasiswa.Transaksi, fmt.Sprintf("Menambah saldo sebesar Rp%d", jumlah))
		fmt.Printf("Saldo berhasil ditambahkan. Saldo saat ini: Rp%d\n", mahasiswa.Saldo)
	} else {
		fmt.Println("Jumlah saldo harus lebih dari 0.")
	}
}

func TarikSaldo(mahasiswa *Mahasiswa) {
	var jumlah int
	fmt.Println("=== Tarik Saldo ===")
	fmt.Print("Masukkan jumlah saldo yang ingin ditarik: ")
	fmt.Scan(&jumlah)

	if jumlah <= 0 {
		fmt.Println("Tidak dapat menarik saldo dengan jumlah 0 atau negatif.")
	} else if jumlah > mahasiswa.Saldo {
		fmt.Println("Saldo tidak cukup untuk melakukan penarikan.")
	} else {
		mahasiswa.Saldo -= jumlah
		mahasiswa.Transaksi = append(mahasiswa.Transaksi, fmt.Sprintf("Menarik saldo sebesar Rp%d", jumlah))
		fmt.Printf("Penarikan berhasil. Saldo saat ini: Rp%d\n", mahasiswa.Saldo)
	}
}

func HistoryTransaksi(mahasiswa Mahasiswa) {
	fmt.Println("=== Riwayat Transaksi ===")
	if len(mahasiswa.Transaksi) == 0 {
		fmt.Println("Tidak ada transaksi.")
	} else {
		for _, transaksi := range mahasiswa.Transaksi {
			fmt.Println(transaksi)
		}
	}
}

func main() {

	mahasiswa1 := Mahasiswa{
		Username:  "Fathimah",
		NPM:       "2406423805",
		Saldo:     3500000,
		Transaksi: []string{},
	}

	var username, password string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password (NPM): ")
	fmt.Scan(&password)

	if username != mahasiswa1.Username || password != mahasiswa1.NPM {
		fmt.Println("Username atau password salah. Program berhenti.")
		os.Exit(1)
	}

	for {
		var pilihan int
		fmt.Println("\n=== Menu Utama ATM ===")
		fmt.Println("1. Lihat Informasi Akun")
		fmt.Println("2. Lihat Saldo")
		fmt.Println("3. Tambahkan Saldo")
		fmt.Println("4. Tarik Saldo")
		fmt.Println("5. Riwayat Transaksi")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			LihatAkun(mahasiswa1)
		case 2:
			LihatSaldo(mahasiswa1)
		case 3:
			TambahSaldo(&mahasiswa1)
		case 4:
			TarikSaldo(&mahasiswa1)
		case 5:
			HistoryTransaksi(mahasiswa1)
		case 6:
			fmt.Println("Keluar dari program. Terima kasih!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
