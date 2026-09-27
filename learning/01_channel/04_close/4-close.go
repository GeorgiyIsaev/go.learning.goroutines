package main

import "fmt"

func main() {
	{ // Закрытие канала
		var ch = make(chan int, 1)

		ch <- 42

		close(ch)

		value, ok := <-ch
		fmt.Println("Value:", value, "IsOpen:", ok)

		value, ok = <-ch
		fmt.Println("Value:", value, "IsOpen:", ok)

		value, ok = <-ch
		fmt.Println("Value:", value, "IsOpen:", ok)

		//close(ch) // Закрытие уже закрытого канала вызывает панику
		//ch <- 42 // Запись в закрытый канал тоже вызывает панику
	}
}
