package main

// 14

import "fmt"

/*
Tipe Data Array
Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
Daya tampung Array tidak bisa bertambah setelah Array dibuat
Array berdasarkan index di array dimulai dari 0
*/

func main() {
	var names [3]string // Variable names mempunyai 3 kumpulan data
	names[0] = "Rizky"  // Mengisikan indexs 0 names
	names[1] = "Faisal" // Mengisikan indexs 1 names
	names[2] = "Rafi"   // Mengisikan indexs 2 names
	//names[3] = "Joko"	// Invalid array index '3' (out of bounds for the 3-element array)

	fmt.Println(names[0]) // data ke-1
	fmt.Println(names[1]) // data ke-2
	fmt.Println(names[2]) // data ke-3
	//fmt.Println(names[3]) // Invalid array index '3' (out of bounds for the 3-element array)

	// Membuat Array Langsung
	// Di Go-Lang juga bisa membuat Array secara langsung saat deklarasi variable
	var values = [3]int{
		90, 80, 20,
	}
	fmt.Println(values)

	/*
		Function Array
		Operasi						Keterangan
		len(array)					Untuk mendapatkan panjang Array
		array[index]				Mendapat data di posisi index
		array[index] = value		Mengubah data di posisi index
	*/
	var values2 = [...]int{
		90, 80, 70,
	}
	var lagi [10]string
	fmt.Println(values2)
	fmt.Println(len(values2)) // Panjang Array 3
	values[0] = 200
	fmt.Println(values2)
	fmt.Println(lagi)

}
