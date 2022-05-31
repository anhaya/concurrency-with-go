package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func main() {
	version := os.Args[1:][0]
	var wgExecuteRequest = func(wg *sync.WaitGroup) {
		defer wg.Done()
		executeRequest(version)
	}
	var wg sync.WaitGroup
	for i := 0; i < 99; i++ {
		wg.Add(1)
		go wgExecuteRequest(&wg)
	}
	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Finished.")
	fmt.Printf("Count is: %s", executeRequest(version))
	fmt.Println()
}

func executeRequest(version string) string {
	httpClient := http.Client{}

	resp, _ := httpClient.Get(fmt.Sprintf("http://127.0.0.1:8080/%s/count", version))
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
