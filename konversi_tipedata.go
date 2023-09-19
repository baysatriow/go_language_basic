package main

import "fmt"

func main() {
	// Konversi tipedata terkadang di butuhkan untuk mengkonversi dari tipedata satu ketipedata lainnya

	// Contoh
	var nilai8 int8 = 127
	var nilai16 int16 = int16(nilai8)
	fmt.Println(nilai8)
	fmt.Println(nilai16)

	// Konversi byte dalam karakter menjadi sebuah karakter
	var nama = "Anto Subarjo"
	fmt.Println(nama)

	// Cari karakter Pertama Dari nama
	var cari byte = (nama[0])
	var huruf string = string(cari)
	fmt.Println("Belum di Konversi ke string", cari)
	fmt.Println("Sudah di Konversi dari byte menjadi string", huruf)
}
