package main

import "log"

/*
	Base nenhum
*/

//

type logSimples struct {
	msg []string
	err error
}

func main() {
	mensagem := logSimples{[]string{"funciona", "log simples", "muito simples"}, nil}

	log.Println(mensagem)
}
