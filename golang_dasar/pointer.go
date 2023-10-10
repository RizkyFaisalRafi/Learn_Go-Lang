package main

import "fmt"

/**
Pointer
Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada.
Sederhananya, dengan kemampuan Pointer, kita bisa membuat passing data by reference.
Untuk mrnggunakannya memakai keyword &.

Operator &
Untuk membuat sebuah variabel dengan nilai pointer ke variabel yang lain,
kita bisa menggunakan operator & diikuti dengan nama variabelnya.


Passing By Value
Secara default di GoLang semua variabel itu di passing data by value, bukan by reference.
Artinya, jika kita mengirim sebuah variabel ke dalam function, method atau variabel lain,
sebenarnya yang dikirim adalah duplikasi valuenya.


Operator *
Saat kita mengubah variabel pointer, maka yang berubah hanya variabel tersebut.
Semua variabel yang mengacu ke data yang sama tidak akan berubah.
Jika kita ingin mengubah seluruh variabel yang mengacu ke data tersebut, kita bisa menggunakan operator *.
Contoh


Function new
Sebelumnya untuk membuat pointer dengan menggunakan operator &
Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal.
*/

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // address1 (Passing by Value) / &address1 (Passing by Reference)
	address3 := &address1

	address2.City = "Tangerang"

	//address2 = &Address{
	//	City:     "Jakarta",
	//	Province: "DKI Jakarta",
	//	Country:  "Indonesia",
	//}

	*address2 = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	fmt.Println(address4) // Kosong
	address4.City = "Bali"
	fmt.Println(address4) // Bali
}
