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
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialForLoop(10))
	fmt.Println(factorialRecursive(10))
}
