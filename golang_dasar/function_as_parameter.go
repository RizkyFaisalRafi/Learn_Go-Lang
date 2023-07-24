package main

// 28

import "fmt"

/*
Function tidak hanya bisa kita simpan didalam variable sebagai value
Namun juga bisa kita gunakan sebagai parameter untuk function lain

Function Type Declarations
Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function
sebagai parameter
*/

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Kucing" {
		return "..."
	} else {
		return name
	}
}

type Filter func(string) string

func sayHelloWithFilter2(name string, filter2 Filter) {
	fmt.Println("hello", filter2(name))
}

func spamFilter2(name string) string {
	if name == "Kucing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Rafi", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Kucing", filter)

	// Function Type Declarations
	sayHelloWithFilter2("Rafi", spamFilter2)

	filter2 := spamFilter2
	sayHelloWithFilter2("Kucing", filter2)
}
