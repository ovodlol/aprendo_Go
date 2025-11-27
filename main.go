package main

/*  Aprendo Go:
Só minhas anotações e práticas sobre golang dividido em capitulos
*/

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print(`
	  Aprendo Go

	Anotações, práticas e muito mais,
	para eu dominar Go :3

	`)

	file, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Print("o README.md não tá aqui!!!")
	}

	fmt.Println(string(file))
}
