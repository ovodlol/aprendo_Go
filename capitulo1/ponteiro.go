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

func logado(msg *string) {
	*msg = fmt.Sprintf("%v aprendo_Go: %s\n", time.Now(), "oi mundo")
}

func main() {
	logado(&mensagem) 

	fmt.Println(mensagem)

	logado(&mensagem)

	logarei(&mensagem)
}

func logarei(msg *string) {
	println(*msg)
}
