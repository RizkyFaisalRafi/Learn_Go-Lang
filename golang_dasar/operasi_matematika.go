package main

// 11

import "fmt"

/*
Operasi Matematika
Operator:
+  ->  Penjumlahan
-  ->  Pengurangan
*  ->  Perkalian
/  ->  Pembagian
%  ->  Sisa Bagi / Modulus

Augmented Assignment
Operasi Matematika			Augmented Assignment
a = a + 10							a += 10
a = a - 10							a -= 10
a = a * 10							a *= 10
a = a / 10							a /= 10
a = a % 10							a %= 10

Unary Operator
Operator					   Keterangan
++								a = a + 1
--								a = a - 1
-								Negative
+								Positive
!								Negasi / Boolean Kebalikan
*/

func main() {
	fmt.Println("Operator")
	var a = 10
	var b = 10

	// Penambahan
	var c = a + b
	fmt.Println(c)

	// Perkalian
	c = a * b
	fmt.Println(c)

	// Pengurangan
	c = a - b
	fmt.Println(c)

	// Pembagian
	c = a / b
	fmt.Println(c)

	// Modulus
	c = a % b
	fmt.Println(c)

	fmt.Println("\nAugmented Assignment")
	a += 10
	fmt.Println(a)

	a -= 10
	fmt.Println(a)

	a *= 10
	fmt.Println(a)

	a /= 10
	fmt.Println(a)

	a %= 10
	fmt.Println(a)

	fmt.Println("\nUnary Operator")
	var i int
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
	var negative = -100
	var Positive = +100 // Sama ajah kaya 100 tanpa kata kunci +
	fmt.Println(negative)
	fmt.Println(Positive)
}
