package main

import "fmt"

/**
Type Assertions
Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data menjadi tipe data yang diinginkan
Konversi tipe data
Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong


Type Assertions Menggunakan Switch
Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
Jika panic dan tidak ter-recover, maka otomatis program akan mati/terhenti
Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
*/

func random() interface{} {
	return 1
}

func main() {
	var result = random()
	// Konversi secara paksa menggunakan type assertions
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	// panic: interface conversion: interface {} is string, not int
	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	// Type Assertions Menggunakan Switch
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}

}
