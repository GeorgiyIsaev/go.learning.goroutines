package main

import (
	"fmt"
	"sync"
	"time"
)

func workerC(in <-chan string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range in {
		fmt.Println("IN:", msg)

		out <- msg + " (обработано)"

		time.Sleep(time.Second) // имитируем работу
	}

	fmt.Println("Горутина завершила работу.")
}

func main() {
	// Out
	out := make(chan string)

	go func() {
		for msg := range out {
			fmt.Println("OUT:", msg)
		}
	}()

	// In
	in := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)
	go workerC(in, out, &wg)

	// Отправляем сообщения в горутину
	in <- "Привет"
	in <- "Как дела?"
	in <- "Пока"

	close(in) // закрываем канал, чтобы завершить горутину

	wg.Wait() // ждём, пока все горутины завершатся

	close(out) // безопасно закрываем канал out, уверенные в том что в него никто не пишет

	fmt.Println("Программа завершена.")
}
