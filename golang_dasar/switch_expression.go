package main

// 18

import "fmt"

/*
Switch Expression
Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
Switch expression sangat sederhana dibandingkan if
Biasanya switch  expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

Switch dengan Short Statement
Sama dengan If, Switch juga mendukung short statement sebelum variable yang akan di cek kondisinya

Switch Tanpa Kondisi
Kondisi di switch expression tidak wajib
Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case nya
*/
func main() {
	name := "Jokowi"

	switch name {
	case "Eko":
		fmt.Println("Hello Pak Eko")
	case "Joko":
		fmt.Println("Hello Pak Joko")
	case "Bagas":
		fmt.Println("Hello Pak Bagas")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length2 > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
