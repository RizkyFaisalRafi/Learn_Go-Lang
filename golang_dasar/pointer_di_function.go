package main

import "fmt"

/**
Pointer di Function
Saat kita membuat parameter di function, secara default adalah pass by value,
artinya data akan di copy lalu dikirim ke function tersebut.
Oleh karena itu, jika kita mengubah data di dalama function, data yang aslinya tidak akan pernah berubah.
Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah.
Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut,
Untuk melakukan ini, kita juga bisa menggunakan pointer di function.
Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya.
*/

type Addresss struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address Addresss) {
	address.Country = "Indonesia"
}

func ChangeAddressToIndonesia2(address *Addresss) {
	address.Country = "Indonesia"
}

func main() {
	address := Addresss{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	ChangeAddressToIndonesia(address)

	fmt.Println(address) // Tidak berubah

	// ==============================================
	// Pointer di Function
	address2 := Addresss{
		City:     "Subang",
		Province: "Jawa Timur",
		Country:  "",
	}

	ChangeAddressToIndonesia2(&address2)

	fmt.Println(address2) // berubah
}
