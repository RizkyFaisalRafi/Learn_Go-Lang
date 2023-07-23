package main

// 19

import (
	"fmt"
)

/*
For
Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan/looping
Salah satu fitur perulangan adalah for loops

For dengan Statement
Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa tambahkan di for
Init statement, yaitu statement sebelum for di eksekusi
Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan

For Range
For bisa digunakan untuk melakukan iterasi terhadap semua data collection
Data collection contohnya Array, Slice dan Map
*/

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke-", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke-", counter)
	}

	names := []string{"Rizky", "Faisal", "Rafi"} // Slice
	// Cara Manual
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Cara Cepat dan efisien
	for index, name := range names {
		fmt.Println("index", index, " = ", name)
	}

	// _ berfungsi agar tidak error karena di golang apabila ada variable yang tidak dipakai maka akan error
	for _, name := range names {
		fmt.Println(name)
	}

	// Map key, value || slice atau array index, value
	for key, value := range names {
		fmt.Println(key, " = ", value)

	}
}
