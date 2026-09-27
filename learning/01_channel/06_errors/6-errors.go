package main

import "fmt"

func main() {
	{ // Работа с неинициализированным каналом
		var ch chan int
		//ch <- 1 // fatal error: all goroutines are asleep - deadlock!
		fmt.Println("Канал из блока 1:", ch)
	}

	{ // Дедлок - никто не читает
		var ch = make(chan int, 1)

		ch <- 1
		//ch <- 1 // fatal error: all goroutines are asleep - deadlock!
		fmt.Println("Канал из блока 2:", ch)
	}

	{ // Дедлок - никто не пишет
		var ch = make(chan int, 1)

		//<-ch // fatal error: all goroutines are asleep - deadlock!
		fmt.Println("Канал из блока 3:", ch)
	}

	{ // Закрытие закрытого канала
		var ch = make(chan int, 1)
		close(ch)
		//close(ch) // panic: close of closed channel
		fmt.Println("Канал из блока 4:", ch)
	}

	{ // Запись в закрытый канал
		var ch = make(chan int, 1)
		close(ch)
		//ch <- 1 // panic: send on closed channel
		fmt.Println("Канал из блока 5:", ch)
	}

	{ // Чтение из закрытого канала - OK!
		var ch = make(chan int, 1)
		close(ch)

		_, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed")
		}
		fmt.Println("Канал из блока 6:", ch)
	}

	{ // For range закрытого канала - OK!
		var ch = make(chan int, 1)

		ch <- 1

		close(ch)

		for value := range ch {
			fmt.Println("Прочитано из канала:", value)
		}

		fmt.Println("Прочитали всё из канала и вышли из цикла")
		fmt.Println("Канал из блока 7:", ch)
	}

	{ // Select закрытого канала - OK!
		var ch = make(chan int, 1)
		close(ch)

		select {
		case value := <-ch:
			fmt.Println("Прочитано из канала (zero value):", value)
		default:
			fmt.Println("Канал закрыт")
		}
		fmt.Println("Канал из блока 8:", ch)
	}
	fmt.Println("END MAIN")
}
