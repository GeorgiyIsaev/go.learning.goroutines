package main

//паттерн graceful shutdown
//— корректное завершение программы по сигналу.
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	tasks := make(chan int, 20)
	stop := make(chan struct{})

	// Server HTTP
	serverDone := make(chan struct{})

	go func() {
	FOR:
		for {
			select {
			case <-stop:
				break FOR
			case tasks <- 42:
			}
		}

		close(serverDone)
	}()

	// UseCase
	usecaseDone := make(chan struct{})

	go func() {
	FOR:
		for {
			select {
			case <-stop:
				break FOR
			case task := <-tasks:
				fmt.Println("Value:", task)
				time.Sleep(time.Second)
			}
		}

		close(usecaseDone)
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal:", <-wait)

	close(stop)

	<-serverDone
	fmt.Println("Server done")

	<-usecaseDone
	fmt.Println("UseCase done")
}
