package main

import (
	"fmt"
	"time"
)

func bubbleSort(arr []int) {
	inicio := time.Now()
	n := len(arr)

	for i := range n {
		for j := range n - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fim := time.Since(inicio)
	fmt.Printf("o normal demorou %v\n", fim)
}

func bubbleSortOt(arr []int) {
	inicio := time.Now()
	n := len(arr)

	for i := range n {
		ainda_pode_mudar := false
		for j := range n - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				ainda_pode_mudar = true
			}
		}

		if !ainda_pode_mudar {
			break
		}
	}
	fim := time.Since(inicio)
	fmt.Printf("o otmizado demorou %v\n", fim)
}

func main() {
	arr := []int{10, 9, 23, 1, 54, 2, 23, 34, 13, 43, 7, 5}
	arrOt := arr
	fmt.Printf("antes: %v\n", arr)

	bubbleSort(arr)
	bubbleSortOt(arrOt)

	fmt.Printf("com o normal depois: %v\n", arr)
	fmt.Printf("com o otimizado depois: %v\n", arrOt)
}
