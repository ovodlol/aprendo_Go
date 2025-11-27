package main

import (
	"fmt"
	"time"
)

func insertionSort(arr []int) {
	inicio := time.Now()

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key

	}
	fim := time.Since(inicio)
	fmt.Printf("ele demorou %v\n", fim)
}

func main() {
	arr := []int{10, 9, 23, 1, 54, 2, 23, 34, 13, 43, 7, 5}
	fmt.Printf("antes: %v\n", arr)

	insertionSort(arr)

	fmt.Printf("depois: %v\n", arr)
}
