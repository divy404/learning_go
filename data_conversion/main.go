package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println(num)
	fmt.Printf("%T",num)

	var data float64 = float64(num);
	fmt.Println("Type of data is: ", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is ", str)
	fmt.Printf("Type of str is %T\n", str)

	number_string := "1234"
	number_int,_ := strconv.Atoi(number_string)
	fmt.Println("str is:",number_int)
	fmt.Printf("Type of str is %T\n",number_int)

	num_string := "3.14"
	number_float,_ := strconv.ParseFloat(num_string,64)
	fmt.Println("number_float is ",number_float)
	fmt.Printf("type of number_float is %T\n",number_float)

}
