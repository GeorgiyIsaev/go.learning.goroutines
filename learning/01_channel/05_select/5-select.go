package main

import "fmt"

func main() {
	{ // Select - конструкция для работы с каналами
		ch1 := make(chan int, 10)
		ch2 := make(chan int, 10)
		ch3 := make(chan int, 10)

		ch1 <- 1
		ch2 <- 2
		ch3 <- 3

		var value int

		select {
		case value = <-ch1:
		case value = <-ch2:
		case value = <-ch3:
		}

		fmt.Println(value) // Случайное число: 1, 2 или 3
	}

	{ // Default
		ch1 := make(chan int, 10)
		ch2 := make(chan int, 10)

		var value int

		select {
		case value = <-ch1:
		case value = <-ch2:
		default:
		}

		fmt.Println("Default case:", value) // 0
	}

	// default нужен просто чтобы выйти из селекта если все кейсы заблокированы.
	// Чтобы упростить код, default оставляют пустой, а логику пишут ниже.
	tasks := make(chan int, 10)
	stop := make(chan int)

	{ // Вот это
		select {
		case <-stop:
			break
		default:
			tasks <- 42 // мы тут можем заблокироваться и не узнать что канал stop закрыт
		}
	}

	{ // И вот это, одно и тоже
		select {
		case <-stop:
			break
		default:
		}
		//но в любом случаи отправим число в канал
		tasks <- 42 // мы тут можем заблокироваться и не узнать что канал stop закрыт
	}
}
