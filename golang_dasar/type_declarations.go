package main

// 10

/*
Type Declarations
Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada,
dengan tujuan agar lebih mudah dimengerti
*/

import "fmt"

func main() {
	type NoKTP string
	type married bool

	var ktpFaisal NoKTP = "123456"
	var marriedStatus married = false

	fmt.Println(ktpFaisal)
	fmt.Println(NoKTP("987654"))

	fmt.Println(marriedStatus)
}
