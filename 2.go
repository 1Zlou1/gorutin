package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		n := 1
		for {
			select {
			case <-sigCh:
				fmt.Println("Получен сигнал прерывания. Выход из программы.")
				os.Exit(0)
			default:
				fmt.Printf("Квадрат %d: %d\n", n, n*n)
				n++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	for {
		select {
		case <-sigCh:
			fmt.Println("Получен сигнал прерывания. Выход из программы.")
			os.Exit(0)
		}
	}
}
