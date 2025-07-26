package main

import "fmt"

func main() {
	fmt.Println("leetsss goo arrayyyyy")
	var name[5] string
	name[0] = "i"
	name[1] = "love"
	name[2] = "pizza"
	fmt.Println(name)

	var num = [5]int{1,2,3,4}
	fmt.Println("Number is:",num)
	fmt.Println("Lenght of numbers is:",len(num))

	fmt.Println("value of name at 2 index is",name[2])
	// go lang sets default value of that data type in array which we are using 
	var pookie[3]string
	fmt.Printf("our pookie is empty:%q\n",pookie)


}
