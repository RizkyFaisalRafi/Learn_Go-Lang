package main

// 15

import "fmt"

/*
Tipe Data Slice
Tipe data Slice adalah potongan dari data Array
Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

Detail Tipe Data Slice
Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
Pointer adalah penunjuk data pertama di array pada slice
Length adalah panjang dari slice, dan
Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

Membuat Slice			Keterangan
array[low:high]			Membuat slice dari array dimulai index low sampai index sebelum high
array[low]				Membuat slice dari array dimulai index low sampai index akhir di array
array[:high]			Membuat slice dari array dimulai index 0 sampai index sebelum high
array[:]				Membuat slice dari array dimulai index 0 sampai index akhir di array

Function Slice
Operasi								Keterangan
len(slice)							Untuk mendapatkan panjang
cap(slice)							Untuk mendapatkan kapasitas
append(slice, data)					Membuat slice baru dengan menambah data ke posisi terakhir slice,
									jika kapasitas sudah penuh, maka akan membuat array baru
make([]TypeData, length, capacity)	Membuat slice baru
copy(destination, source)			Menyalin slice baru dari source ke destination
*/

func main() {
	var months = [...]string{ // ... dipakai jika tidak tau kapasitasnya berapa
		"Januari",   // 0
		"Februari",  // 1
		"Maret",     // 2
		"Apil",      // 3
		"Mei",       // 4
		"Juni",      // 5
		"Juli",      // 6
		"Agustus",   // 7
		"September", // 8
		"Oktober",   // 9
		"November",  // 10
		"Desember",  // 11
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // Length
	fmt.Println(cap(slice1)) // Capacity

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Faisal")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	copySlice2 := make([]string, 1, cap(newSlice))
	copy(copySlice2, newSlice)
	fmt.Println(copySlice2)

	// Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array, melainkan slice
	iniArray := [...]int{1, 2, 3, 4, 5, 6}    // bisa ... atau langsung angka didalam kurung siku
	iniSlice := []int{1, 2, 3, 4, 5, 6, 7, 8} // Kurung Sikunya Kosong yaitu Slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
