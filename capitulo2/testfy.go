package main

/*
   Base no "https://github.com/bahlo/go-styleguide?tab=readme-ov-file#testing"
   e no "https://gobyexample.com/testing-and-benchmarking"

   Muito importante saber fazer testes
*/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	valor := conta(2, 4, 3)
	espera := 4
  assert.Equal(t, espera, valor)
}

func conta(x, y, z int) int {
	return x * y / z
}
