package main

/*
	uma simples conta até dez com goroutimes
	onde um conta de 1 até 10 outro faz a conta
	1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10
	que é 55 :3
*/

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("vou contar até 10")
		for i := range 10 {
			fmt.Printf("conto %v\n", i+1)
			time.Sleep(time.Second)
		}
		fmt.Println("terminei")
		wg.Done()
	}()

	go func() {
		var conta int
		for i := range 10 {
			conta = conta + (i + 1)
			time.Sleep(time.Second)
		}
		fmt.Printf("tudo é %v", conta)
		wg.Done()
	}()

	wg.Wait()
}
