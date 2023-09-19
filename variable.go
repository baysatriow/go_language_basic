package main

import "fmt"

func main() {

	// Variable berfungsi untuk menyimpan sebuah data, dapat berupa integer, floating, string dll.

	var nama_lengkap = "Bayu Satrio Wibowo"
	var tanggal_lahir = "12 Maret 2005"
	var tempat_lahir = "Pringsewu"
	var umur = 18

	fmt.Println("Halo Nama Saya", nama_lengkap, ", Saya Lahir di", tempat_lahir, "Pada tanggal", tanggal_lahir, ", Umur Saya saat ini adalah", umur)

	// Nilai dari variable dapat berubah ubah sesuah dengan alur program
	umur += 2

	fmt.Println("Saat ini tahun 2025 maka Umur saya adalah ", umur)
}
