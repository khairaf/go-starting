package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i:=0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}	
	//wg.Done()
}


func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	say("there you go")
	//time.Sleep(1000*time.Millisecond)
	

	wg.Add(1)
	go say("Buzz")
	wg.Wait()
}

/*package main

import "fmt"

func foo() {
  defer fmt.Println("Done!")
  fmt.Println("Doing some stuff, who knows what?")
}

func main() {
  foo()
}*/