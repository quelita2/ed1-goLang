// MergeSort: dividir para conquistar

package main

import ("fmt")

// Função que divide os vetores até estarem com tamanho 1  (seu ultimo estado divisível) e dps executar a conquista sobre eles (DIVISÂO)
func mergeSort(v []int) []int {
    if len(v) > 1{
        // dividir o vetor em 2 vetores: criação de ve e vd
        metade := len(v)/2
        ve := make([]int, metade)
        vd := make([]int, len(v) - metade)
        
        iv := 0
        
        // preenchendo os vetores ve e vd
        for i:=0; i < metade; i++ {
            ve[i] = v[iv]
        }
        
        for i:=0; i < len(v) - metade; i++ {
            vd[i] = v[iv]
        }
        
        // chamada recursiva para dividir o vetor em 2 vetores
        ve = mergeSort(ve)
        vd = mergeSort(vd)
        v = merge(ve, vd)
    }
    return v
}

// Função que mescla os vetores ordenados (CONQUISTA)
// primissa: ve e vd estão ordenados
func merge(ve []int, vd []int) []int {
	// criar vetor ordenado com tam = tam(v1)+tam(v2)
    v := make([]int, len(ve)+len(vd))
    
    // indices que percorrem os vetores 
    iv := 0
    ive := 0
    ivd := 0
    
    // laço comparando elementos de ve e vd, copiando os menores para v
    for ive < len(ve) && ivd < len(vd) {
        if ve[ive] < vd[ivd] {
            v[iv] = ve[ive]
            ive++
        } else {
            v[iv] = vd[ivd]
            ivd++
        }
        iv++
    }

    // laço que tenta copiar os elementos restantes de ve para v
    for ive < len(ve) {
        v[iv] = ve[ive]
        ive++
        iv++
    }

    // laço que tenta copiar os elementos restantes de vd para v
    for ivd < len(vd){
        v[iv] = vd[ivd]
        ivd++
        iv++
    }

    return v
}

func main() {
    v := []int{12, 11, 13, 5, 6, 7}
    fmt.Println("Vetor não ordenado:", v)
    v = mergeSort(v)
    fmt.Println("Vetor ordenado:", v)
}