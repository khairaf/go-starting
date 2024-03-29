package main

import "fmt"

func foo(c chan int, someValue int){
	c <- someValue*5 //add hasil perkalian tersebut ke channel c
}

func main (){
	fooVal := make(chan int)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	v1, v2 := <- fooVal, <- fooVal
	fmt.Println(v1, v2) //15 25
}

// <- adalah channel operator