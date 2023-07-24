package main

import "fmt"

// 24

/*
Function Return Multiple Values
Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya
di function.

Menghiraukan Return Value
Multiple return value wajib ditangkap semua value nya
Jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda _ (garis bawah)
*/

func getFullName() (string, string) {
	return "Faisal", "Rafi"
}

func getFullName2() (string, string, string) {
	return "Rizky", "Faisal", "Rafi"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println(getFullName())

	firstName2, _, _ := getFullName2()
	fmt.Println(firstName2)
}
