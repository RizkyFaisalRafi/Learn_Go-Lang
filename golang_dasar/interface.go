package main

import "fmt"

/**
Interface
Interface adalah tipe data yang abstract, dia tidak memiliki implementasi langsung
Sebuah interface berisikan definisi definisi method
Biasanya interface digunakan sebagai kontrak
Interface bisa dijadikan parameter di function

Sebelumnya di struct bisa bikin data langsung di struct. tapi kalau di interface tidak bisa


Implementasi Interface
Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
Sehingga kita tidak perlu mengimplementasikan interface secara manual
Hal ini agak berbeda dengan bahasa pemrograma lain yang ketika membuat interface, kita harus menyebutkan secara ekplisit
akan menggunakan interface mana

*/

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// Kode program implementasi Interface (1)
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// Kode program implementasi Interface (2)
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var rafi Person
	rafi.Name = "Rafi"
	SayHello(rafi) // Output: Hello Rafi

	animal := Animal{Name: "Kucing"}
	SayHello(animal) // Output: Hello Kucing
}
