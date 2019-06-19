// package main

// //cara deklarasi import lebih dari 1 librari
// //fmt: format package
// import (
// 			"fmt"
// 			"time"
// 			"math"
// 			"math/rand")
// // math/rand artinya mau pake rand package yang berada dalam math package (package dalam package)
// func foo(){
// 	fmt.Println("foo!")
// }

// func main() {
// 	fmt.Println("Pertama-tama!")
// 	fmt.Println("Akar 4 =",math.Sqrt(4)) //yang didalam kurung di fungsi Sqrt disebut parameter
// 	fmt.Println("Random angka dari 0 - 100 adalah ", rand.Intn(10))
// 	foo()
// 	fmt.Println("Sekarang jam:", time.Now())
	
// }

//fungsi main selalu running otomatis tanpa invoke seperti di js, fungsi selain main barulah musti di invoke didalam main itu sendiri.
//godoc fmt Println untuk baca dokumentasi tentang fmt Println di terminal

//Tipe data musti dispecify. variable dan parameter, dan return-an dari sebuah function musti ditulis tipenya. 

/*
list of the Go language types:
-bool
-string
-int  int8  int16  int32  int64
-uint uint8 uint16 uint32 uint64 uintptr
-byte // alias for uint8
-rune // alias for int32
     // represents a Unicode code point
-float32 float64
-complex64 complex128
*/

/*
package main

import "fmt"

//func add (x float64, y float64) float64 {
func add (x, y float64) float64 {
	return x+y
}

func greeting(a,b string)(string, string){
	return a,b
}

func main() {
	// var num1 float64 = 2.2
	// var num2 float64 = 3.3
	var num1,num2 = 2.3, 4.2
	greet1, greet2 := "What's", " Up?!"
	fmt.Println(add(num1, num2))
	fmt.Print(greeting(greet1, greet2))

}
*/
/*
When you define, if you don't assign a value, strings default to "", booleans default to false and numerical types default to 0.

	var a int = 62
	var b float64 = float64(a)
	
	type inference works in Go:

	var x float32
	y := x // y is float32 type

*/

//pointer
package main

import "fmt"

func main(){
	x := 15
	fmt.Println(x)
	a := &x //&x points ke memori addrss x
	fmt.Println(a) //bakal ngeprint memoru addresnya: 0xc000052080
	fmt.Println(*a) //bakal ngeprint value dari variable tersebut :15
	*a = 5
	fmt.Println(x) //5
	
}