package main

import "fmt"

/*
Defer, Panic, Recover
Penjelasan Defer
Defer function adalah function yang bisa kita jadwalkan untuk di eksekusi setelah sebuah funtion selesai di eksekusi
Defer function akan selalu di eksekusi walaupun terjadi error di function yang di eksekusi

Penjelasan Panic
Panic function adalah function yang bisa digunakan untuk menghentikan program
Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

Penjelasan Recover
Recover adalah function yang bisa digunakan untuk menangkap data panic
Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
Recover yang benar disimpan di Defer function
*/

// Kode Defer
func logging() {
	fmt.Println("Selesai memanggil function")
}
func runApplication(value int) {
	// defer logging() memberitahu golang bahwa setelah selesai memanggil func runApplication
	// maka akan diakhiri memanggil defer logging
	defer logging() // Biasakan menggunakan defer diawal
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

// ==========================================================

// Kode Panic dan Recover
func endApp() {
	// NOTE: Jika uncomment akan memakai recover / sebaliknya jika comment tidak memakai recover
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message: ", message)

	}
	fmt.Println("Aplikasi Selesai")

}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error Brother")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	//runApplication(0) // panic: runtime error: integer divide by zero

	//runApp(true)  // panic: Error Brother
	runApp(true)

	fmt.Println("Rafi")

}
