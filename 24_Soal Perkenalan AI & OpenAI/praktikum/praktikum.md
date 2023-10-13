pada implementasi AI ini saya menggunakan OpenAi yang dimana saya membuat suatu condition pada system.

System :
kamu sebagai generate code programming yang akan menggenerate inputan kedalam code bahasa golang

jika inputan itu meminta selain bahasa golang berikan respon "Maaf Saya Tidak Bisa Memberikan Bahasa Selain Golang"

Disini saya mencoba untuk menggunakan conditions yang salah.
Input :
berikan code rumus segitiga menggunakan bahasa java

Output :
Maaf, saya hanya bisa memberikan contoh kode dalam bahasa Golang. Jika Anda memerlukan contoh kode dalam bahasa Golang, saya siap membantu.

Lalu saya mencoba untuk menggunakan condition yang benar.
Input :
berikan code rumus segitiga mengguankan bahasa golang

Output :
Tentu, berikut adalah contoh kode untuk menghitung luas segitiga menggunakan bahasa Golang:

```go
package main

import (
	"fmt"
)

func main() {
	var alas float64
	var tinggi float64

	fmt.Print("Masukkan panjang alas: ")
	fmt.Scanln(&alas)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&tinggi)

	luas := 0.5 * alas * tinggi

	fmt.Println("Luas segitiga adalah:", luas)
}
```

Kode di atas akan meminta input dari pengguna untuk panjang alas dan tinggi segitiga, kemudian menghitung luas segitiga dan mencetak hasilnya.

Dengan Begitu system yang saya buat berfungsi dengan baik.
