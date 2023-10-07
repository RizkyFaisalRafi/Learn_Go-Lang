package main

import "fmt"

// 26

/*
Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
Apa bedanya dengan parameter biasa dengan tipe data Array?
	- Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
	- Jika parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu,
      cukup gunakan tanda koma

Slice Parameter
Kadang ada kasus dimana kita menggunakan Variadic Function, namun memiliki variable berupa slice
Kita bisa menjadikan slice sebagai vararg parameter

*/

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)
	total2 := sumAll(10, 10, 10, 10)
	fmt.Println(total2)

	// Slice Parameter
	numbers := []int{10, 10, 10, 10}
	total2 = sumAll(numbers...)
	fmt.Println(total2)
}
