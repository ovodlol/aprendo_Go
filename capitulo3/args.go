package main

/*
    Base no meu conhecimento e criatividade(as vozes)
		Go é muito usado em cli então vou aprender a fazer,
		vou usar a biblioteca padrão "flag"(para argumentos e flags)
		e o "fmt"(para mostrar as coisas na tela).
*/

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "batata": 
		fmt.Println("Batata é um legume que é muito saudavel")
	case "banana":
		fmt.Println("Banana é uma fruta boa")
	case "help":
		fmt.Println("tem que mandar assim: ","\n\t<comando> batata","\n\t<comando> banana")
	default:
		fmt.Println("se tiver duvida coloque <comando> help")
	}
}
