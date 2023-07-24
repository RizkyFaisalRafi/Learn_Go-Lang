package main

import "fmt"

// 25

/*
Named Return Values
Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data
return value di function

Namun kita juga bisa membuat variable secara langsung di tipe data return function nya
*/

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Rizky"
	middleName = "Faisal"
	lastName = "Rafi"

	//return firstName, middleName, lastName
	return // Tidak perlu di deklrasikan seperti diatas
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	a, b, c := getCompleteName()
	fmt.Println(getCompleteName())
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
