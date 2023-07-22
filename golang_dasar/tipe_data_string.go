package main

/*
Tipe Data String
String adalah tipe data kumpulan karakter
Jumlah karakter didalam String bisa nol/kosong sampai tidak terhingga
Tipe data String di Go-Lang direpresentasikan dengan kata kunci string
Nilai data String di Go-Lang selalu diawali dengan karakter " (Petik Dua) dan diakhiri " (Petik Dua)

Function untuk String
len("string") => Menghitung jumlah karakter di String
"string"[number indexs] => Mengambil karakter pada posisi yang ditentukan
*/

import "fmt"

func main() {
	fmt.Println("Rizky")
	fmt.Println("Faisal")
	fmt.Println("Rafi")

	fmt.Println(len("Joko"))
	fmt.Println("Kurniansyah"[0]) // Hasilnya 75 karena 75 adalah representasi bytenya si K
	fmt.Println("Zero"[1])        // Hasilnya 101 karena 101 adalah representasi bytenya si e
}
