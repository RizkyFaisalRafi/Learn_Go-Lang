package main

/*
Konversi Tipe Data
Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain.

Hati hati jika konversi dari valuenya besar ke tipe data lebih kecil ukuran tidak cukup,
bisa menyebabkan perubahan valuenya / integer overflow.
*/

import "fmt"

func main() {
	var nilai32 int32 = 32768
	// Konversi tipe data dari 32 ke 64
	var nilai64 int64 = int64(nilai32)
	// Konversi tipe data dari 32 ke 16
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Rizky Faisal"
	var e = name[0]         // Mengambil index ke 0 dari value variable name namun hasilnya byte
	var eString = string(e) // Konversi byte sebelumnya menjadi string

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
