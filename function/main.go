package main

import "fmt"

func simpleFunction() {
	fmt.Println("this is our func")
}

func add(a,b int) (int) {
	return a+b
}
// the int mentioned nex to  the fun paramter is for the type of output we want 

func main() {
	fmt.Println("I am happy learning goooo hehe")
	simpleFunction()

	ans:= add(1,2)
	println(ans)
}
