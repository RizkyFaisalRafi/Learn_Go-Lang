package main

import "fmt"

/**
Struct
Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
Struct biasanya representasi data dalam program aplikasi yang kita buat
Data di struct disimpan di dalam field
Sederhananya struct adalah kumpulan dari field

Membuat Data Struct
Struct adalah template data atau prototype data
Struct tidak bisa langsung digunakan
Namun kita bisa membuat data/object dari struct yang telah kita buat

Struct Literals
Sebelumnya kita telah membuat data dari struct,
namun sebenarnya ada banyak cara yang bisa kita gunkan untuk membuat data dari struct

Struct Method
Struct adalah tipe data seperti tipe data lainnya, bisa digunakan sebagai parameter untuk function
Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan akan sebuah struct memiliki function
Method adalah funtion
*/

type CustomerPremium struct {
	Name, Address string
	Age           int
}

func main() {
	var faisal CustomerPremium
	faisal.Name = "Rizky Faisal Rafi"
	faisal.Address = "Pondok Alam Permai"
	faisal.Age = 21

	fmt.Println(faisal)         // {Rizky Faisal Rafi Purati 21}
	fmt.Println(faisal.Name)    // Rizky Faisal Rafi
	fmt.Println(faisal.Address) // Pondok Alam Permai
	fmt.Println(faisal.Age)     // 21

	// Program Struct Literals
	joko := CustomerPremium{
		Name:    "Joko Darwin",
		Address: "Keroncong Permai",
		Age:     22,
	}
	fmt.Println(joko) // {Joko Darwin Keroncong Permai 22}

	budi := CustomerPremium{"Budi Sentosa", "Total Persada", 25} // Tidak direkomendasikan, selengkapnya tonton lgi

	fmt.Println(budi) // {Budi Sentosa Total Persada 25}

}
