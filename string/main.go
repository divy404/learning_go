package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,banana,orange"
	parts := strings.Split(data,",")
	fmt.Println(parts)

	str:="one two three four five two two five two"
	count := strings.Count(str,"two")
	fmt.Println("count:",count)

	str1:= "     dharamrajj gaaanduuuu      hehe       "
	trimmed:= strings.TrimSpace(str1)
	fmt.Println("trimmed:",trimmed)
	// TrimSpace just checks for space in start and end 

	str2:="divi"
	str3:="jain"
	result:=strings.Join([]string{str2,str3},"")
	fmt.Println("result:",result)



}