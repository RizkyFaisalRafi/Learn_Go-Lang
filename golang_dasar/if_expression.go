package main

// 17

import "fmt"

/*
IF EXPRESSION
If adalah salah satu kata kunci yang digunakan untuk percabangan
Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
Hampir di semua bahasa pemrograman mendukung if expression

ELSE
Blok if akan dieksekusi ketika kondisi if bernilai true
Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
Hal ini bisa dilakukan menggunakan else expression

ELSE IF
Kadang dalam If, kita butuh membuat beberapa kondisi
Kasus seperti ini, kita bisa menggunakan Else If expression

If dengan Short Statement
If mendukung short statement sebelum kondisi
Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi
*/
func main() {
	name := "Doni"

	if name == "Jokan" {
		fmt.Println("Hello Jokan")
	} else if name == "Doni" {
		fmt.Println("Hello Doni")
	} else {
		fmt.Println("Hello Anonymous")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
