package main

import (
	"errors"
	"fmt"
)

type Result struct {
	Error    error
	response string
}

func main() {

	call1 := make(chan Result)
	call2 := make(chan Result)

	go func() {
		fmt.Println("Calling first API")
		call1 <- Result{Error: errors.New("any error")}
	}()

	go func() {
		fmt.Println("Calling second API")
		call2 <- Result{response: "concurrency in go"}
	}()

	result1 := <-call1
	result2 := <-call2

	if result1.Error != nil {
		fmt.Println("Error in first call")
	}
	if result2.Error != nil {
		fmt.Println("Error in second call")
	}
}
