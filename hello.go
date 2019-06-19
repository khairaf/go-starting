package main

//cara deklarasi import lebih dari 1 librari
//fmt: format package
import (
			"fmt"
			"time"
			"math"
			"math/rand")
// math/rand artinya mau pake rand package yang berada dalam math package (package dalam package)
func foo(){
	fmt.Println("foo!")
}

func main() {
	fmt.Println("Pertama-tama!")
	fmt.Println("Akar 4 =",math.Sqrt(4)) //yang didalam kurung di fungsi Sqrt disebut parameter
	fmt.Println("Random angka dari 0 - 100 adalah ", rand.Intn(10))
	foo()
	fmt.Println("Sekarang jam:", time.Now())
	
}

//fungsi main selalu running otomatis tanpa invoke seperti di js, fungsi selain main barulah musti di invoke didalam main itu sendiri.
//godoc fmt Println untuk baca dokumentasi tentang fmt Println di terminal