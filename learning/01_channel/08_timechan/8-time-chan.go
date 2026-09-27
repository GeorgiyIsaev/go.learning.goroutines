package main

import (
	"fmt"
	"time"
)

func main() {
	//{ // Ticker
	//for range time.Tick(time.Second) {
	//	fmt.Println("Tick")
	//}
	//	//Бесконечный цикл
	//}

	{ // Timer
		timer := time.NewTimer(5 * time.Second)

		fmt.Println("Сработает через 1 сек")
		<-timer.C
		fmt.Println("Сработало!")

		timer.Stop()
	}

	{ // After - то же самое, что и Timer, но короче
		fmt.Println("Сработает через 1 сек")
		<-time.After(time.Second)
		fmt.Println("Сработало!")
	}

	{ // Timeout
		ch := make(chan struct{})

		select {
		case <-ch:
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout")
		}
	}
}
