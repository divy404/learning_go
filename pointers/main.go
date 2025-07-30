package main

import "fmt"

func modifyValueByRef (num *int) {
	*num = *num*20

}
func main() {
	// var num int
	num := 2

	// var ptr *int

	ptr := &num

	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println("Accesing data through pointer:",*ptr)
	// the *int shows the typo of data which we have assigned to that block 


	// var pointer *int
	// if pointer == nil {
	// 	fmt.Println("Pointer is not assigned.")
	// }

	value := 10
	modifyValueByRef(&value)
	fmt.Println("Value contains: ", value)
}