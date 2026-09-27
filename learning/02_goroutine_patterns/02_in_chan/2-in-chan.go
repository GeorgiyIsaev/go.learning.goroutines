package main

//паттерн «один отправитель — один получатель»
import (
	"fmt"
	"sync"
	"time"
)

func workerB(in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range in {
		fmt.Println("Получено сообщение:", msg)

		time.Sleep(time.Second) // имитируем работу
	}

	fmt.Println("Горутина завершила работу.")
}

func main() {
	in := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)
	go workerB(in, &wg)

	// Отправляем сообщения в горутину
	in <- "Привет"
	in <- "Как дела?"
	in <- "Пока"

	close(in) // закрываем канал, чтобы завершить горутину

	wg.Wait() // ждём, пока все горутины завершатся

	fmt.Println("Программа завершена.")
}
