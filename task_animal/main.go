package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// doAction выполняет одно действие для животного:
// выводит начало, три шага с паузой 1 секунда и завершение.
// После завершения отправляет код действия в канал results
// и обновляет переменную lastAnimal (имя последнего выполнившего действие).
func doAction(animal string, code rune, name string, results chan<- string, lastAnimal *string, mu *sync.Mutex) {
	// Начало действия
	fmt.Printf("%c %s начало %s\n", code, animal, name)

	// Три шага выполнения (имитация процесса)
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%c %s %s\n", code, animal, name)
	}

	// Завершение действия
	fmt.Printf("%c %s закончило %s\n", code, animal, name)

	// Отправляем код действия в канал результатов
	results <- string(code)

	// Запоминаем, какое животное выполнило это действие последним
	mu.Lock()
	*lastAnimal = animal
	mu.Unlock()
}

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Канал для очереди животных (буфер 3)
	animals := make(chan string, 3)
	animals <- "Кошка"
	animals <- "Собака"
	animals <- "Попугай"
	close(animals) // закрываем, чтобы горутины могли завершиться после чтения

	totalActions := 10 // общее количество действий

	// Случайное распределение 10 действий между тремя животными
	// каждое животное получит хотя бы 1 действие
	var counts [3]int
	for {
		a := rand.Intn(8) + 1 // 1..8
		b := rand.Intn(8) + 1
		c := totalActions - a - b
		if c >= 1 && c <= 8 {
			counts[0] = a
			counts[1] = b
			counts[2] = c
			break
		}
	}

	var wg sync.WaitGroup
	results := make(chan string, totalActions) // канал для сбора кодов действий
	var lastAnimal string
	var mu sync.Mutex

	// Запускаем три горутины – каждая обрабатывает одно животное
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			// Получаем животное из очереди (канал)
			animal := <-animals
			count := counts[idx]

			// Выполняем заданное количество действий для этого животного
			for j := 0; j < count; j++ {
				// Случайный выбор действия
				action := rand.Intn(3)
				var code rune
				var name string
				switch action {
				case 0:
					code = 'E'
					name = "ест"
				case 1:
					code = 'S'
					name = "спит"
				case 2:
					code = 'P'
					name = "играет"
				}
				doAction(animal, code, name, results, &lastAnimal, &mu)
			}
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	close(results)

	// Собираем коды действий в порядке их завершения
	var letters []string
	for letter := range results {
		letters = append(letters, letter)
	}

	// Выводим результат: последовательность кодов и имя последнего животного
	fmt.Print("Результат: ")
	for _, l := range letters {
		fmt.Print(l + " ")
	}
	fmt.Println(lastAnimal)
}
