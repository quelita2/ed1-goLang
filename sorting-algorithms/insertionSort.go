package main

import "fmt"

func insertionSort(v []int) {
    n := len(v)
    
    for i := 1; i < n; i++ {
        key := v[i]
        j := i - 1
        
        // Mova os elementos da parte ordenada da lista para a direita
        // até encontrar a posição correta para inserir o elemento atual.
        for j >= 0 && v[j] > key {
            v[j+1] = v[j]
            j--
        }
        
        v[j+1] = key
    }
}

func main() {
    v := []int{64, 25, 12, 22, 11}
    insertionSort(v)
    fmt.Println("Lista ordenada:", v)
}
