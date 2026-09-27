package main

import (
	"fmt"
	"unsafe"
)

func main() {
	{ // Канал - кольцевая очередь с мьютексом внутри (безопасный для многопоточности)
		var ch chan int // nil-канал который не имеет смысла

		fmt.Println("ch == nil:", ch == nil) // true

		//ch <- 1 // panic: deadlock
		//<-ch // panic: deadlock

		ch = make(chan int, 10) // Создание канала с буфером в 10 элементов

		// Запись
		ch <- 42
		ch <- 42
		ch <- 42

		// Чтение
		fmt.Println("Прочитали из канала:", <-ch) // 42

		fmt.Println("Длина канала:", len(ch))                     // 2
		fmt.Println("Размер канала на стеке:", unsafe.Sizeof(ch)) // 8
	}
}
