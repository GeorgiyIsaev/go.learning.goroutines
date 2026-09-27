package main

import (
	"fmt"
	"strconv"
	"time"
)

func worker(worker int, ch chan<- string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		ch <- "Hello from worker " + strconv.Itoa(worker)
	}
}

func main() {
	tasks := make(chan string, 10)

	for i := range 10 {
		go worker(i, tasks)
	}

	for value := range tasks {
		fmt.Println(value)
	}
	//! Никогда не завершиться
}
