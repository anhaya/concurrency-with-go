package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := generate("Hello")
	c2 := generate("There")

	fanin := make(chan string)
	go func() {
		for {
			select {
			case str1 := <-c1:
				fanin <- str1
			case str2 := <-c2:
				fanin <- str2
			}

		}
	}()

	go func() {
		for {
			fmt.Println(<-fanin)
		}
	}()

}

func generate(data string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- data
			time.Sleep(time.Duration(100 * time.Millisecond))

		}
	}()
	return channel
}
