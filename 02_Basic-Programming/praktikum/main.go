package main

import (
	"fmt"
	"strings"
)

//Function Prioritas 1 No1
func Luas (Atas, Bawah, Tinggi int)int{
	return (Atas + Bawah)* Tinggi /2
}

//Function Prioritas 1 No2
func Bilangan (angka int)string{
	if angka %2 == 0 {
		return "Genap"
	}
	return "Ganjil"
}

//Function Prioritas 1 No3
func Grade (nilai int)string{

	if (nilai >= 80) && (nilai <= 100) {
		return  "A"
	}else if (nilai >= 65) && (nilai <= 79){
		return  "B"
	}else if (nilai >= 50) && (nilai <= 64){
		return  "C"
	}else if (nilai >= 35) && (nilai <= 49){
		return  "D"
	}else if (nilai >= 0) && (nilai <= 34) {
		return  "E"
	}else{
		return  "Nilai Invalid"
	}
	}

//Function Eksplorasi
func polindrome (kata string)bool{
	kata = strings.ToLower(kata)
	teks := len(kata)

	for i := 0; i < teks/2; i++{
		if kata[i] != kata[teks -1-i] {
			return false
		}
	} 
	return true
}

func reverse(kata string) string {
	runes := []rune(kata)
	teks := len(runes)

	for i := 0; i < teks/2; i++ {
		runes[i], runes[teks-1-i] = runes[teks-1-i], runes[i]
	}

	return string(runes)
}

func main(){
	for{
		fmt.Println("\nTugas Basic Programming ")
		fmt.Println("1. Prioritas 1")
		fmt.Println("2. Prioritas 2")
		fmt.Println("3. Eksplorasi")

		var pilih int
		fmt.Printf("Masukkan Pilihan : ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
			for{
				fmt.Println("\nTugas Prioritas 1 ")
				fmt.Println("1. Soal 1")
				fmt.Println("2. Soal 2")
				fmt.Println("3. Soal 3")
				fmt.Println("4. Soal 4")
		
				var pilih int
				fmt.Printf("Masukkan Pilihan : ")
				fmt.Scanln(&pilih)
		
		//Tugas Prioritas 1 No1
			if pilih == 1{
				fmt.Println("\nPERHITUNGAN LUAS TRAPESIUM")
				var Atas, Bawah, Tinggi int
				fmt.Print("Masukkan Nilai Atas : ")
				fmt.Scanln(&Atas)
				fmt.Print("Masukkan Nilai Bawah : ")
				fmt.Scanln(&Bawah)
				fmt.Print("Masukkan Nilai Tinggi : ")
				fmt.Scanln(&Tinggi)
				Hasil := Luas(Atas, Bawah, Tinggi)
				fmt.Printf("Luas Trapesium :%d \n\n", Hasil)
				break
		
		//Tugas Prioritas 1 No2
			}else if pilih == 2 { 
				fmt.Println("\nMENENTUKAN BILANGAN GANJIL GENAP")
				var angka int
				fmt.Print("Masukkan Angka : ")
				fmt.Scanln(&angka)
				Hasil := Bilangan(angka)
				fmt.Printf("%d Adalah Bilangan %s \n\n", angka, Hasil)
				break
		
		//Tugas Prioritas 1 No3
			}else if pilih == 3 { 
				fmt.Println("\nPEMBERIAN GRADE PADA NILAI")
				var nilai int
				fmt.Print("Masukkan Nilai :")
				fmt.Scanln(&nilai)
				Hasil := Grade (nilai)
				fmt.Printf("Grade nilai : %s \n\n",Hasil)
				break
		
		//Tugas Prioritas 1 No4
			}else if pilih == 4 { 
				fmt.Println("\nPEMBERIAN FIZZ BUZZ")
			for i := 1; i <= 100; i++ {
				if i%3 == 0 && i%5 == 0 {
					fmt.Print("Fizz Buzz ")
				} else if i%3 == 0 {
					fmt.Print("Fizz ")
				} else if i%5 == 0 {
					fmt.Print("Buzz ")
				} else {
					fmt.Printf("%d ", i)
				}
			}
				fmt.Println()
				break
			}else { 
				fmt.Println("Pilihan Tidak Valid")
				break
			}	
			}
		break 
		}else if pilih == 2 {
			fmt.Println("\nTugas Prioritas 2 ")
			fmt.Println("1. Soal 1")
			fmt.Println("2. Soal 2")

			var pilih int
			fmt.Printf("Masukkan Pilihan : ")
			fmt.Scanln(&pilih)

		//Tugas Prioritas 2 no1
		if pilih == 1 {
			fmt.Println("\nMencetak Segitiga Asterik ")
			N := 5
			for i := 1; i <= N; i++{
				for j := 1 ; j <= N - i; j++ {
					fmt.Print(" ")
				}
				for k := 1; k <= i; k++ {
					fmt.Print(" *")
				}
				fmt.Println(" ")
				
			}
		//Tugas Prioritas 2 No2
		}else if pilih == 2 {
			fmt.Println("\nMencetak Faktor Bilangan ")
			var input int
			fmt.Print("Masukkan angka: ")
			fmt.Scan(&input)

			fmt.Printf("Faktor-faktor dari %d adalah: ", input)
			for i := 1; i <= input; i++ {
				if input%i == 0 {
				fmt.Print(i, " ")
				}
			}
			fmt.Println()
		} 
		break
		
		//Tugas Eksplorasi
		}else {
			fmt.Println("\nTugas Eksplorasi ")
			var input string
			fmt.Print("Masukan Kata : ")
			fmt.Scanln(&input)

			reversed := reverse(input)
			fmt.Printf("Captured : %s\n", reversed)

			if polindrome(input){
				fmt.Printf("%s Adalah Polindrome. \n", input)
			}else{
				fmt.Printf("%s Bukan Polindrome. \n", input)
			}

		} 
		break

	}
}