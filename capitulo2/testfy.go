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

func TestAdd(t *testing.T) {
	conta  := 2 + 2
	espera := 4
  assert.Equal(t, espera, conta)
}

