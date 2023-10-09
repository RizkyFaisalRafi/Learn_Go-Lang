package main

import "fmt"

/**
Struct Method / function struct function
Struct adalah tipe data seperti tipe data lainnya, bisa digunakan sebagai parameter untuk function
Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan akan sebuah struct memiliki function
Method adalah funtion
*/

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My name is", customer.Name)
}

func (customer2 Customer) sayHi(name string) {
	fmt.Println("Hi", name, "My name is", customer2.Name)
}

func main() {
	dudi := Customer{Name: "Dudi Sanjaya"}
	dudi.sayHello()

	var joko Customer
	joko.Name = "Joko"
	joko.Age = 22
	joko.Address = "Tangerang"

	doni := Customer{Name: "Doni Jonathan"}
	doni.sayHi("Jokowi")
}
