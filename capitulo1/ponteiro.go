package main

/*
		Base nenhum só meus incriveis conhecimentos do basico
		Ponteiros, eles apontão para uma váriavel
		isso deixa o programa mais leve
*/

import (
	"time"
	"fmt"
)

var mensagem string

func logarei(msg *string) {
	ano := time.Now().Format("01/02 15:04:05")
	*msg = fmt.Sprintf("%v aprendo_Go: %s\n", ano,  "oi mundo")
}

func main() {
	logarei(&mensagem) 

	fmt.Println(mensagem)
}
