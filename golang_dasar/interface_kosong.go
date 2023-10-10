package main

import "fmt"

/**
Interface Kosong
Golang bukanlah bahasa pemrograman berorientasi objek
Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap
sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
Contoh di Java ada java.lang.Object

Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
Interface Kosong adalah interface yang tidak memiliki deklarasi method satupun,
hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	} else {
		return "Hello Other"
	}
}

func main() {
	var data interface{} = Ups(34)
	fmt.Println(data)
}
