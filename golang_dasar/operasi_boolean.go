package main

// 13

import "fmt"

/*
Operator			Keterangan
&&					AND / DAN
||					OR /  ATAU
!					NOT / NEGASI
*/

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus) // False
}
