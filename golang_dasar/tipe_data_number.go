package main

// 4

/*
2 Jenis Tipe Data Number yaitu Integer(Bilangan Bulat) dan Floating Point (Desimal)

Tipe Data Integer (1)
- int8 [-128 ] -> [127]
- int16 [-32768] -> [32767]
- int32 [-2147483648] -> [2147483647]
- int64 [-9223372036854775808] -> [9223372036854775807]

Tipe Data Unsigned Integer(Tidak punya nilai minimum negatif dan nilai maksimumnya ditambah, misal 127 + 128 = 255) (2)
uint8 [0] -> [255]
uint16 [0] -> [65535]
uint32 [0] -> [4294967295]
uint64 [0] -> [18446744073709551615]

Tipe Data Floating Point
https://docs.google.com/presentation/d/1QNFV9kjV4TfN-FVFLT6-8Urq2MmadAmgc1puk-YE5Fs/edit#slide=id.g76f41bdaac_0_67

Penggunaan tipe data di golang itu Case Sensitive
*/

import "fmt"

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 1)
	fmt.Println("Tiga Koma Lima = ", 3.5)
}
