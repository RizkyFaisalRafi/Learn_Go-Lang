package main

// 3

// fmt merupakan singkatan dari format,
// berguna untuk melakukan formating input dan output.
// memungkinkan Anda untuk mencetak dan membaca data dalam berbagai format.
// fmt.Print(), fmt.Println(), fmt.Printf() (Mencetak teks atau nilai ke output standar (console))
// fmt.Scan(), fmt.Scanf() (Membaca input dari pengguna)

/*
Operator ":=" adalah operator deklarasi dan inisialisasi dalam bahasa pemrograman Go (Golang).
Ini digunakan untuk membuat dan menginisialisasi variabel baru dalam satu langkah.
Dengan operator ":=", Go akan mendeteksi tipe data variabel secara otomatis berdasarkan nilai yang diberikan.
*/

import "fmt"

func main() {
	name := "Faisal" // Mendeklarasikan variabel "name" dan menginisialisasinya dengan nilai "Faisal"
	usia := 21       // Mendeklarasikan variabel "usia" dan menginisialisasinya dengan nilai 21
	fmt.Println("Hello World")
	fmt.Println(name, usia)
}
