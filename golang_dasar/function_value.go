package main

import "fmt"

// 27

/*
Function Value
Function adalah first class citizen
Function juga merupakan tipe data, dan bisa disimpan di dalam variable

*/

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Faisal"))

	goodBye2 := getGoodBye("Rafi")
	fmt.Println(goodBye2)
}
