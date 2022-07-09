package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	recordsC := readFile("file.txt")

	start := time.Now()
	for val := range removeInvalid(toUpperCase(recordsC)) {
		fmt.Printf("%v\n", val)
	}
	fmt.Printf("done after %v", time.Since(start))
}

func readFile(fileName string) <-chan []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan []string)

	go func() {
		defer f.Close()
		defer close(ch)

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			s := strings.Split(scanner.Text(), ",")
			time.Sleep(1 * time.Second)
			ch <- s
		}
	}()

	return ch
}

func toUpperCase(strC <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		defer close(ch)
		for val := range strC {
			val[0] = strings.ToUpper(val[0])
			val[1] = strings.ToUpper(val[1])
			time.Sleep(1 * time.Second)
			ch <- val
		}
	}()

	return ch
}

// Remove "invalid" organization/repository. Instead make an GitHub API request
// Ill just remove if its an odd line, after 3 seconds
func removeInvalid(strC <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		defer close(ch)
		i := -1
		for val := range strC {
			i++
			if i%2 != 0 {
				continue
			}
			time.Sleep(1 * time.Second)
			ch <- val
		}
	}()

	return ch
}
