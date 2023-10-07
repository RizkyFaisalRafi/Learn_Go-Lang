package main

import "fmt"

// 30

/*
Recursive Function adalah function yang memanggil dirinya sendiri

Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function
lebih mudah dibandingkan tidak menggunakan recursive function

Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah factorial
*/

func factorialForLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 { // Kondisi berhenti
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialForLoop(5))   // 5 * 4 * 3 * 2 * 1 = 120
	fmt.Println(factorialRecursive(5)) // 5 * 4 * 3 * 2 * 1 = 120
}
