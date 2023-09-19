package main

import "fmt"

//Tipe Data Integer (1) Start From -
// int8, int16, int32, int64

func main() {
	fmt.Println("Ini Bilangan Integer Tipe 1 =", -125)

	// Tipe Data Integer (2) Start From 0
	// uint8, uint16, uint32, uint64

	fmt.Println("Ini Bilangan Integer Tipe 2 =", 3000)

	// Tipe Data Floating Point
	// float16, float32, complex64, complex128

	fmt.Println("Ini Bilangan Floating Point =", 0.0015)

	// Alias
	//  byte(uint8), rune(int32), int(32/64), uint(32/64)

	fmt.Println("Ini Contoh Bilangan Menggunakan Tipe Deklarasi Alias")

	var alias_1 byte = 255

	fmt.Println("Alias byte pengganti uint8 dengan nilai minimum 0 dan maksimum 255 seperti contoh = ", alias_1)

	var alias_2 rune = 2147483647

	fmt.Println("Alias rune penggati int32 dengan nilai minimum -2147483648 dan maksimum 2147483647 seperti contoh = ", alias_2)

	var alias_3 int = -999999999

	fmt.Println("Alias int dengan minimum int32 dan maksimum int64 (Mengikuti Base Sistem Operasi yang digunakan) seperti contoh = ", alias_3)

	var alias_4 uint = 999999999

	fmt.Println("Alias uint pengganti deklarasi dengan minimum uint32 dan maksimum uint64 (Mengikuti Base Sistem Operasi yang digunakan) seperti contoh = ", alias_4)
}
