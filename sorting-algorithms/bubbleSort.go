package main

import "fmt"

func bubbleSort(v []int) {
    n := len(v)
    // O loop externo controla o número de iterações, realizando n-1 iterações, pois o ultimo elemento já permanecerá ao final do vetor
    for i := 0; i < n-1; i++ {
        // O loop interno compara pares adjacentes e realiza as trocas, se necessário
        for j := 0; j < n-i-1; j++ {
            if v[j] > v[j+1] {
                // Troca os elementos se estiverem fora de ordem
                v[j], v[j+1] = v[j+1], v[j]
            }
        }
    }
}

func main() {
    v := []int{64, 25, 12, 22, 11}
    bubbleSort(v)
    fmt.Println("Lista ordenada:", v)
}
