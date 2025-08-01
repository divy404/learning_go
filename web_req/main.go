package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web services")
	 res,err:=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	 if err != nil {
		fmt.Println("error getting GET response", err)
		return
	 }
	 defer res.Body.Close()

	 data,err := io.ReadAll(res.Body)
	 if err != nil {
		fmt.Println("Error reading response", err)
		return
	 }
	 fmt.Println("response: ", string(data))

}