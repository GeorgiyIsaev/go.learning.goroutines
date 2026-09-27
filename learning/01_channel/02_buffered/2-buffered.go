package main

import "fmt"

func main() {
	{ // Буферизированный канал
		var ch = make(chan int, 3)

		ch <- 1
		ch <- 2
		ch <- 3

		value := <-ch // Читаем канал и присваиваем значение в переменную

		fmt.Println(value) // 1

		<-ch // Читаем канал и отбрасываем значение

		value = <-ch

		fmt.Println(value) // 3

	}
}
