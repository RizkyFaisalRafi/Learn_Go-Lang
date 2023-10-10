package main

import "fmt"

/**
Nil
Biasanya didalam bahasa pemrograman lain, object yang belum di inisialisasi maka secara otomatis nilainya adalah null / nil
Berbeda dengan GoLang, di GoLang saat kita buat variabel dengan tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya
Namun di GoLang ada data nil, yaitu data kososng
Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer, channel.

Default value di Golang apabila belum di inisialisasi:
int = 0
boolean = false
string = ""
struct = struct kosong

tapi di golang ada nil bisa digunakan sebagai data kosong

Nil bisa digunakan untuk pengecekan
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	var person map[string]string = NewMap("")

	// Pengecekan Tidak Pakai Nil
	if person["name"] == "" {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

	// Pengecekan Pakai Nil
	if person == nil {
		fmt.Println("Data Kosong Bro nil")
	} else {
		fmt.Println(person)
	}
}
