//BubbleSort implementa e demonstra o referido algoritmo de ordenação em um array de inteiros.
package main

import "fmt"

func bubbleSort(n []int) []int {
	for i := range n {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
	return n
}

func main() {
	unsorted := []int{1, 7, 8, 5, 9, 3, 2}
	fmt.Println("Array inicial:", unsorted)
	bubbleSort(unsorted)
	fmt.Println("Array ordenado:", unsorted)
}
