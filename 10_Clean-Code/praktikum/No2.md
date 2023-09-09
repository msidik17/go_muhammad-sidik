package main

type kendaraan struct {
totalroda int
kecepatanperjam int
}

type mobil struct {
kendaraan
}

func (m \*mobil) berjalan() {
m.tambahkecepatan(10)
}

func (m \*mobil) tambahkecepatan(kecepatanbaru int) {
m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

func main() {
mobilcepat := mobil{}
mobilcepat.berjalan()
mobilcepat.berjalan()
mobilcepat.berjalan()

    mobillamban := mobil{}
    mobillamban.berjalan()

}

pada code diatas ada beberapa hal yang dapat dilakukan untuk memudahkan membaca code ini seperti :

- menggunakan huruf kapital pada variable yang memiliki lebih dari satu kata contohnya 'totalroda' pada
  struktur data kendaraan lebih baik menggunakan 'TotalRoda'
- menggunakan penamaan yang lebih deskriftif untuk fungsi variable
- menggunakan operator yang tepat seperti 'm.kecepatanperjam = m.kecepatanperjam + kecepatanbaru' pada function 'm\* mobil' lebih baik menggunakan 'm.KecepatanPerJam += kecepatanBaru' untuk mempersingkat code agar lebih mudah untuk dibaca.

Dan berikut adalah code yang telah saya revisi untuk mempermudah dalam pembacaan code

package main

type Kendaraan struct {
TotalRoda int
KecepatanPerJam int
}

type Car struct {
Kendaraan
}

func (m \*Car) berjalan() {
m.tambahKecepatan(10)
}

func (m \*Car) tambahKecepatan(kecepatanBaru int) {
m.KecepatanPerJam += kecepatanBaru
}

func main() {
mobilCepat := Car{}
mobilCepat.berjalan()
mobilCepat.berjalan()
mobilCepat.berjalan()

    mobilLamban := Car{}
    mobilLamban.berjalan()

}
