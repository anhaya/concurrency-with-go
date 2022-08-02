package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	Month int
}

func main() {
	done := make(chan bool)
	defer close(done)

	start := time.Now()

	datas := make(chan Data)
	//Gets the datas from the whole year  (12 months)
	go func() {
		for i := 1; i <= 12; i++ {
			monthData := getDataByMonth(i)
			datas <- monthData
		}
		close(datas)
	}()

	workers := make([]<-chan int, 4)
	for i := 0; i < 4; i++ {
		workers[i] = buildMetric(done, datas, i)
	}

	numMonths := 0
	for range merge(done, workers...) {
		numMonths++
	}

	fmt.Printf("Took %fs to calculate %d months\n", time.Since(start).Seconds(), numMonths)
}

// Get datas to build metrics, it takes 1 seconds
func getDataByMonth(month int) Data {
	time.Sleep(1 * time.Second)
	return Data{Month: month}
}

func buildMetric(done <-chan bool, datas <-chan Data, workerId int) <-chan int {
	metrics := make(chan int)
	go func() {
		for data := range datas {
			select {
			case <-done:
				return
			case metrics <- data.Month:
				time.Sleep(4 * time.Second)
				fmt.Printf("Worker #%d: Calculating metrics for month %d, took 4s to calc\n", workerId, data.Month)
			}
		}
		close(metrics)
	}()

	return metrics
}

func merge(done <-chan bool, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	wg.Add(len(channels))
	outgoingMonths := make(chan int)
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case outgoingMonths <- i:
			}
		}
	}
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(outgoingMonths)
	}()
	return outgoingMonths
}
