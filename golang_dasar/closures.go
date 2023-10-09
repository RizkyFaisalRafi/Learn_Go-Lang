package main

import "fmt"

/*
Closures
Closure adalah kemampuan sebuah function berinteraksi dengan data data disekitarnya dalam scope yang sama
Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi
karena jika salah implementasi closure akan bisa secara tidak sengaja mengubah data yang harusnya tidak diubah
*/

func main() {
	named := "Rizky"
	counter := 0

	increment := func() {
		//a := "B"
		fmt.Println("Increment")
		counter++
		named = "Budi" // Named diluar scope yaitu Rizky keubah menjadi Budi
	}
	//a // Tidak bisa akses di scope dalamnya increment

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(named)

}
