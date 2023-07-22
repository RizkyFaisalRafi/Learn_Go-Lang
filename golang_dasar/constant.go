package main

// 8

/*
Constant
Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya.

Deklarasi Multiple Constant
Sama seperti variable, di Go-lang juga kita bisa membuat constant secara banyak lebih dari 1

	const (
		countryIndo = "Indonesia"
		countryArab = "Arab"
	)

*/

import "fmt"

func main() {
	// const walaupun variablenya belum dipakai, dia tidak akan error, beda dengan variable/var.
	const firstName string = "Eko"
	const lastName = "Kurniawan"
	const value = 100

	//	Error karena const tidak dapat diubah valuenya
	//firstName = "Joko"
	//lastName = "Sugandar"

	fmt.Println(firstName)
	//fmt.Println(lastName)
	fmt.Println(value)

	const (
		countryIndo string = "Indonesia"
		countryArab        = "Arab"
		value2             = 1000
	)
}
