package main

import (
	"fmt"
	"time"
)

func main() {
	selectWithLoop()
	fmt.Println("Program has finished")
}

func selectWithLoop() {
	myNameFunc := func(myName chan string, done chan interface{}) { //this func will close the channel in case API error
		fmt.Println("Calling the My Name API with success")
		time.Sleep(2 * time.Second)
		myName <- "Carlos"
	}
	myPetNameFunc := func(myPetName chan string, done chan interface{}) { //this func will close the channel in case API error
		fmt.Println("Calling the My Pet Name API with error")
		time.Sleep(1 * time.Second)
		close(done)
	}
	myCatNameFunc := func(myCatName chan string, done chan interface{}) { //this func will close the channel in case API error
		fmt.Println("Calling the My Cat Name API with success")
		myCatName <- "I do not have a cat"
	}

	myName := make(chan string)
	myPetName := make(chan string)
	myCatName := make(chan string)
	done := make(chan interface{})
	go myNameFunc(myName, done)
	go myPetNameFunc(myPetName, done)
	go myCatNameFunc(myCatName, done)
	for { //if we do not have this for, its impossible to assert that "done" and "myCatName" cases would be executed
		select {
		case name := <-myName:
			fmt.Println(name)
		case name := <-myPetName:
			fmt.Println(name)
		case name := <-myCatName: //Its only falls here, because myPetFunc (the one that closes) run before myNameFunc
			fmt.Println(name)
		case <-done: //myPetFunc fall here
			fmt.Println("channel was closed")
			return
		}
	}

}
