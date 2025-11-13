package main

/*
		Falo um pouco sobre chanels(canais em portugues),
		por que é muito importante.
		Os canais são filas para os gorotines
		falarem entre si 
		
		Exemplo de channels:
			- função com um loop
				- coloca "e " na fila do canal
				- imprime o intex mais 1
			- função main 
				- cria o canal de strings
				- executa a outra função 
				- espera um pouco
				- imprime o canal inteiro
*/

import (
	"fmt"
	"time"
)

func tacaNoCanal(msg chan string) {
	for i := range 3 {
		msg <- "e "
		fmt.Print(i+1, " ")	
	}
}

func main() {
	msg := make(chan string)

	go tacaNoCanal(msg)

	time.Sleep(time.Second * 3)

	fmt.Print(<- msg)
	fmt.Print(<- msg)
	fmt.Print(<- msg)
}
