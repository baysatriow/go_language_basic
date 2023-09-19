package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var age int8 = 30
	fmt.Println(age)

	country := "America"
	fmt.Println(country)

	var (
		angka1 = 20
		angka2 = 10
	)
	fmt.Println("Perhitungan", angka1, " : ", angka2)
	fmt.Println("Jumlah", angka1*angka2)
}
