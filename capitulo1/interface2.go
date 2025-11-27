package main

/*
	Base no "https://github.com/cuongndc9/7_days_of_go/blob/main/day%201/16_interfaces.go"
	Um pouco mais de coisas
*/

import (
	"fmt"
)

type animal struct {
	nome, color string
}

func (a animal) miado() {
	fmt.Printf("meow meow %s\n", a.nome)
}

func (a animal) latido() {
	fmt.Printf("au au %s\n", a.nome)
}

type gato interface {
	miado()
}

type cachorro interface {
	latido()
}

func main() {
	var toby cachorro = animal{"Toby", "preto e branco"}
	var nico gato = animal{"nico", "branco"}

	toby.latido()
	nico.miado()
}
