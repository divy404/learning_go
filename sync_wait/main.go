package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n",i)
	// some task is happening 
	fmt.Printf("worker %d end\n",i)
}

func main() {
	var wg sync.WaitGroup
	for i:=1;i<=3;i++{
		wg.Add(1)   //Increment the WaitGroup counter
		worker(i,&wg);
	}
	wg.Wait()
	fmt.Println("workers task complete")
}