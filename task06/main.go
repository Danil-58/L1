package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func MethodChannels() {
	fmt.Println("Старт функции на каналах")

	ch := make(chan struct{})

	go func() {
		defer close(ch)
		defer fmt.Println("Горутина на каналах завершена")
		for i := 0; i < 3; i++ {
			fmt.Println("Горутина на канале: ", i)
			time.Sleep(time.Second)
		}
	}()
	// Ждем завершающего сигнала
	<-ch
}

func MethodContext() {
	fmt.Println("Старт функции на сontext")
	ctx, cancel := context.WithCancel(context.Background())
	defer fmt.Println("Горутина на context завершена")
	defer cancel()

	go func() {
		for {
			select{
			case <-ctx.Done():
				return
			default:
				fmt.Println("Горутина на context работает")
				time.Sleep(time.Second)
			}
		}
	}()
	fmt.Println("Нажмите Enter чтобы остановить горутину на context")
	fmt.Scanln()
}

func MethodWaitGroup() {
	fmt.Println("Старт функции на WaitGroup")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("Горутина на WaitGroup ", i)
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("Горутина на Wait Group завершена")
}

func main() {
	MethodChannels()
	fmt.Println()

	MethodContext()
	fmt.Println()

	MethodWaitGroup()
	fmt.Println()

	fmt.Println("Программа завершена")
}