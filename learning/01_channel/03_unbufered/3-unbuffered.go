package main

import (
	"fmt"
	"time"
)

func main() {
	{ // Небуферизированный канал
		var ch = make(chan int)

		go func() {
			for {
				value, ok := <-ch
				if !ok {
					fmt.Println("Канал закрыт! Выходим")

					return
				}

				fmt.Println("Получили:", value)

				time.Sleep(time.Second)
			}
		}()

		ch <- 1
		ch <- 2
		ch <- 3

		fmt.Println("Отправили все 3 значения")

		close(ch)

		time.Sleep(3 * time.Second)
	}
}
