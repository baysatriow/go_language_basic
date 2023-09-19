package main

import "fmt"

func main() {
	// Tipe data string dapat menyimpan data dalam bentuk karakter dan bisa bernilai kosong maupun tak terhingga

	var hari string = "Senin"
	var jadwal string = ""

	fmt.Println("Sekarang Hari", hari, "Kemudian", jadwal)

	// Beberapa Function untuk String dapat menghitung jumlah karakter dan byte dari karakter

	var nama_lengkap string = "Bayu Satrio Wibowo"
	fmt.Println(nama_lengkap)
	// Menentukan Jumlah karakter dalam string
	var jumlah_str = len(nama_lengkap)
	fmt.Println(jumlah_str)

	// Menentukan karakter dari Sebuah String
	var huruf string = string(nama_lengkap[0])

	fmt.Println(nama_lengkap[0])
	fmt.Println(huruf)
	fmt.Println("Contoh Penggunaan Function untuk String. Nama saya", nama_lengkap, "Nama Saya terdiri dari dari", jumlah_str, "Karakter (termasuk spasi), dan Inisial dari nama saya adalah", huruf)
}
