package main

/*
	Base vozes da minha cabeça elas são bem
	criativas pode chamar elas de Dorinhas

	Uma struct muito simples chamada comida
	e crio uma variavel comida_da_vovó
	onde Arroz verdadeiro, Feizão verdadeiro
	e Mistura é uma lista de strings:
	Bife, Batata frita e Salada
*/

import (
	"github.com/k0kubun/pp/v3"
)

type comida struct {
	Arroz   bool
	Feizão  bool
	Mistura []string
}

func main() {
	comida_da_vovó := comida{
		Arroz:   true,
		Feizão:  true,
		Mistura: []string{"Bife", "Batata frita", "Salada"},
	}

	pp.Println(comida_da_vovó)
}
