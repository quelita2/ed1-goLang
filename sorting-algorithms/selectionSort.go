package main

import("fmt")

func SelectionSort (v [] int){
    n := len(v)

	// O loop será até n-1, visto que o ultimo elemento já permanecerá ao final do vetor
	for i := 0; i < n-1; i++ {
		menor := i  // inicialmente, o menor valor está no 1° indice pois queremos que o menor valor fique no menor indice
		
		for j := menor+1; j < n; j++{
			if v[j] < v[i]{
				// alterar o valor da posição
				aux := v[i]
				v[i] = v[j]
				v[j] = aux
				// atualizar o indice do menor valor
				menor = j
			}
		}
	}
}

func main(){
	v := []int{64, 25, 12, 22, 11}
	SelectionSort(v)
	fmt.Println("Lista ordenada:", v)
}