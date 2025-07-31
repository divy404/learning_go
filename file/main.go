package main

import (
	"fmt"
	
	"os"
)

func main() {
	// file , err := os.Create("example.txt")
	// if err!= nil {
	// 	fmt.Println("error in creating file")
	// 	return
	// }
	// defer file.Close()

	// content := "Hellooo guysss"
	// _, errors := io.WriteString(file,content)
	// if errors != nil {
	// 	fmt.Println("Error while writing file: ", errors)
	// 	return
	// }
	// // io.WriteString(file,content)

	// fmt.Println("filee created yayaayaya")

	// file,err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("error in reading file",err)
	// 	return
	// }
	// defer file.Close()
	// // buffer to read the content of the file
	// buffer := make([]byte,1024)
	// for {
	// 	n,err:=file.Read(buffer)
	// 	if err == io.EOF{
	// 		break
	// 	}
	// 	if err!= nil {
	// 		fmt.Println("Error while reading the file",err)
	// 		return
	// 	}
	// 	fmt.Println(string(buffer[:n]))
	// }

	content,err:= os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("we got some error",err)
	}
	fmt.Println("our content is:",string(content))
}