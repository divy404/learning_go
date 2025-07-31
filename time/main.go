package main

import (
	"fmt"
	"time"
)
func main() {
	currentTime1:= time.Now()
	fmt.Println(currentTime1)
	fmt.Printf("type of current time: %T\n",currentTime1)

	formatted:= currentTime1.Format("02-01-2006, Monday")
	// always need to enter this date only
	fmt.Println(formatted)

	layout_str:="2006-01-02"
	datestr:="2025-07-31"
	formatted_time,_:=time.Parse(layout_str,datestr)
	fmt.Println(formatted_time)


}