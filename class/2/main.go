package main

import (
	"20241021/class/2/models"
	"fmt"
	"math/rand"
)

func LemparDadu(ch chan<- []models.Dadu, n int) {
	sisiDadu := []int{1, 2, 3, 4, 5, 6}
	var daftarLemparanDadu []models.Dadu

	fmt.Printf("Melempar dadu %d kali\n", n)
	for i := 0; i < n; i++ {
		hasilLemparan := models.InitDadu(sisiDadu[rand.Intn(len(sisiDadu))])
		daftarLemparanDadu = append(daftarLemparanDadu, *hasilLemparan)
		fmt.Printf("lemparan %d : %d ", i+1, models.Cetak(*hasilLemparan))
	}
	fmt.Printf("\n\n*** hasil lemparan dadu sebanyak %dx berhasil disimpan ***\n\n ", n)

	ch <- daftarLemparanDadu
}

func BacaLemparanDadu(ch <-chan []models.Dadu, done chan<- bool) {
	lemparanDadu := <-ch
	for i, dadu := range lemparanDadu {
		fmt.Printf("lemparan %d : %d ", i+1, models.Cetak(dadu))
	}
	fmt.Println()
	done <- true
}

func main() {
	dataCh := make(chan []models.Dadu)
	doneCh := make(chan bool)

	var jumlahLemparan int
	fmt.Print("Jumlah lemparan : ")
	fmt.Scanln(&jumlahLemparan)

	go LemparDadu(dataCh, jumlahLemparan)
	go BacaLemparanDadu(dataCh, doneCh)

	<-doneCh
	fmt.Printf("\nSelesai melempar dadu sebanyak %dx, mencatat dan menampilkan hasilnya", jumlahLemparan)
}
