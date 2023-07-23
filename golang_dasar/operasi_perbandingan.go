package main

// 12

import "fmt"

/*
Operasi Perbandingan
Operasi perbandingan adalah operasi untuk membandingkan dua buah data/nilai
Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
Jika hasil operasinya adalah benar, maka nilainya adalah true
Jika hasil operasinya adalah salah, maka nilainya adalah false

Operator					Keterangan
>							Lebih Dari
<							Kurang Dari
>=							Lebih Dari Sama Dengan
<=							Kurang Dari Sama Dengan
==							Sama Dengan
!=							Tidak Sama Dengan
*/

func main() {
	var name1 = "Faisal"
	var name2 = "Rafi"

	var result bool = name1 == name2
	fmt.Println(result)
}
