package main

import (
	"fmt"
	"time"
)

func generate(data string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- data
			time.Sleep(1000)
		}
	}()

	return channel
}

type Processor struct {
	jobChannel chan string
	done       chan *Worker
	workers    []*Worker
}
type Worker struct {
	name string
}

func (w *Worker) processJob(data string, done chan *Worker) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, w.name)
		time.Sleep(3000)
		done <- w
	}()

}

func GetProcessor() *Processor {
	p := &Processor{
		jobChannel: make(chan string),
		workers:    make([]*Worker, 5),
		done:       make(chan *Worker),
	}
	for i := 0; i < 5; i++ {
		w := &Worker{name: fmt.Sprintf("<Worker - %d>", i)}
		p.workers[i] = w
	}
	p.startProcess()
	return p
}

func (p *Processor) startProcess() {
	go func() {
		for {
			select {
			default:
				if len(p.workers) > 0 {
					w := p.workers[0]
					p.workers = p.workers[1:]
					w.processJob(<-p.jobChannel, p.done)
				}
			case w := <-p.done:
				p.workers = append(p.workers, w)
			}
		}
	}()
}

func (p *Processor) postJob(jobs <-chan string) {
	p.jobChannel <- <-jobs
}

func main() {
	source := generate("data string")
	process := GetProcessor()

	for i := 0; i < 12; i++ {
		process.postJob(source)
	}

}
