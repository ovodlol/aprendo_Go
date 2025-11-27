package main

import (
	"fmt"
	"time"
)

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	meio := len(arr) / 2

	esquerda := mergeSort(arr[:meio])
	direita := mergeSort(arr[meio:])

	return merge(esquerda, direita)
}

func merge(esquerda, direita []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(esquerda) && j < len(direita) {
		if esquerda[i] < direita[j] {
			final = append(final, esquerda[i])
			i++
		} else {
			final = append(final, direita[j])
			j++
		}
	}
	for ; i < len(esquerda); i++ {
		final = append(final, esquerda[i])
	}
	for ; j < len(direita); j++ {
		final = append(final, direita[j])
	}
	return final
}

func main() {
	arr := []int{10, 9, 23, 1, 54, 2, 23, 34, 13, 43, 7, 5}
	fmt.Printf("antes: %v\n", arr)

	inicio := time.Now()

	final := mergeSort(arr)

	fim := time.Since(inicio)
	fmt.Printf("ele demorou %v\n", fim)

	fmt.Printf("com o normal depois: %v\n", final)
}
