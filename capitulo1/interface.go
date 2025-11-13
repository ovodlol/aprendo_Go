package main
/*
		Base no https://gobyexample.com/interfaces
		um exemplo simples só para aprender interfaces
		talves melhore o exemplo no futuro colocando
		em prática o effective Go.

		Exemplo de interface:
			- interface geometria que tem duas funções
			  - area que calcula a area do retangulo
				- perimetro que calcula o perimetro do retangulo
			- função fala, que mostra a altura, a largura, a area e o perimetro
*/

import (
	"fmt"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type retan struct {
	width, height float64
}

func (r retan) area() float64{
	return r.width * r.height
}

func (r retan) perimetro() float64{
	return 2*r.width + 2*r.height
}

func fala(g geometria) {
	fmt.Printf("retangulo:   %v\n", g)
	fmt.Printf("area:        %v\n", g.area())
	fmt.Printf("perimetro:   %v\n", g.perimetro())
}

func main() {
	r := retan{width: 8, height: 4}

	fala(r)
}
