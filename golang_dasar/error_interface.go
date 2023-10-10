package main

import (
	"errors"
	"fmt"
)

/**
Error Interface
Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interfacenya adalah error.

Membuat Error
Untuk membuat error, kita tidak perlu manual.
GoLang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (package akan
dibahas secara detail di materi tersendiri).
*/

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak bisa 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError.Error())

	hasil, err := Pembagi(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
