package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Im a developer and forget to put timeout in the next call")
	go func() {
		fmt.Println("Calling an api that do some proccess and returns void")
		for {

		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("My program will finish and kill my goroutine")
	//but in a cenario where my program didn't finish yet, the goroutine would still alive, causing leakeage, I mean
	//its consuming more memory than its necessary
}
