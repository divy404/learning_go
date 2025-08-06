package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hellooooo")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("sayHello function excecuted")
}
func sayHi() {
	fmt.Println("hii")
}

func main() {
	fmt.Println("go_routines")
	go sayHello()
	sayHi()
	time.Sleep(2000*time.Microsecond)

}