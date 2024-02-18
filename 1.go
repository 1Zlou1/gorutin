package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	inputCh := make(chan int)
	squareCh := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			var input string
			fmt.Scanln(&input)
			if input == "стоп" {
				close(inputCh)
				return
			}
			num, err := strconv.Atoi(input)
			fmt.Println("Ввод:", num)
			if err == nil {
				inputCh <- num
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range inputCh {
			square := num * num
			fmt.Println("Квадрат:", square)
			squareCh <- square
		}
		close(squareCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for square := range squareCh {
			result := square * 2
			fmt.Println("Произведение:", result)

		}
	}()

	wg.Wait()
}
