package main

import "fmt"

// 23

/*
Function bisa mengembalikan data

Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut

Jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam function nya kita harus
mengembalikan data

Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya
*/

func getHello(name string) string {
	//return "hello" + name
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Rafi")
	fmt.Println(result)
	fmt.Println(getHello("Faisal"))
	fmt.Println(getHello(""))

}
