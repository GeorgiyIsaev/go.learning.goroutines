package main

import (
	"fmt"
	"sync"
	"time"
)

func workerD(out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, msg := range []string{"Привет", "Как дела?", "Пока"} {
		out <- msg
		fmt.Println("Отправил msg:", msg)
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

	var wg sync.WaitGroup

	wg.Add(1)
	go workerD(out, &wg)

	wg.Wait() // ждём, пока все горутины завершатся

	close(out) // безопасно закрываем канал out, уверенные в том что в него никто не пишет

	fmt.Println("Программа завершена.")
}
