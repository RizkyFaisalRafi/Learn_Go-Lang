package main

/*
Variable
Variable adalah tempat untuk menyimpan data
Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
Di Go-Lang variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda beda jenis,
kita harus membuat beberapa variable
Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya
Variable bisa diubah nilainya

Saat membuat variable, maka wajib menyebutkan tipe data variable tersebut
Namun jika langsung menginisialisasikan data/value pada variablenya, maka tidak wajib menyebutkan tipe data variablenya

Di Go-Lang, kata kunci var saat membuat variable tidaklah wajib.
Asalkan saat membuat variable kita langsung menginisialisasi datanya
Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan
data pada variable tersebut
name := "Faisal"
fmt.Println(name)

Deklarasi Multiple Variable
Di Go-Lang kita bisa membuat variable secara sekaligus banyak
Code yang dibuat akan lebih bagus dan mudah dibaca
var (
	firstName = "Rafi"
	lastName = "Ahmad"
)
*/

import "fmt"

func main() {
	var name1 string = "Rizky Faisal Rafi"
	var name5 = "Rizky Faisal Rafi"
	var name2 string
	name2 = "Eko Kurniawan"

	fmt.Println(name1)
	fmt.Println(name2)

	// Mengubah variable name2
	name2 = "Khanedy"
	fmt.Println(name2)
	fmt.Println(name5)

	var age1 = 30
	var age2 int8 = 30
	fmt.Println(age1)
	fmt.Println(age2)

	country := "Indonesia" // := hanya untuk deklarasi awal
	fmt.Println(country)
	country = "America" // Kalau ada perubahan value tetap pakai =
	fmt.Println(country)

	// Deklarasi Multiple Variable
	var (
		firstName = "Rafi"
		lastName  = "Ahmad"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
