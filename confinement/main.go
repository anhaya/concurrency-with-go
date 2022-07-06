package main

import "fmt"

func main() {
	//adhoc()
	//lexical()
}

func adhoc() {
	data := make([]int, 4)
	loopData := func(handleData chan<- int) { //Its just a function to put the data into a channel
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}
	handleData := make(chan int)
	go loopData(handleData)
	for num := range handleData {
		fmt.Println(num)
	}
	//So, as you can see, the variable "data" should be used just to populate the channel, to handle those informations
	//we should use the channel, but we have a problema here because the variable "data" is available
	//to be accessed from anywhere. By convention, it should be out of loopData function.
}

func lexical() {
	loopData := func(handleData chan<- int) { //Its just a function to put the data into a channel
		data := make([]int, 4)
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}
	handleData := make(chan int)
	go loopData(handleData)
	for num := range handleData {
		fmt.Println(num)
	}
	//So, as you can see, we solved our problem we had in adhoc function.
}
