package main

/*
		Base no "https://gobyexample.com/goroutines"
		só um exemplo de goroutimes muito simples,
		não gostei tão simple que nem deu para entender nada
		sinto que não vi o potencial do goroutines
*/

import (
	"fmt"
	"time"
)

func conta(do string) {
	for i := range 5 {
		fmt.Printf("%s:  %v\n", do, i)
	}
}

func main() {
	conta("bob")

	go conta("maria")

	go func(msg string) {
		time.Sleep(time.Second)
		fmt.Println(msg)
	}("termineido")

	time.Sleep(time.Second)
	fmt.Println("cabo")
}
