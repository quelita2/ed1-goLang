package main

import (
	"fmt"
)

type List interface {
	Add(value int)
	AddPos(value int, pos int)
	Update(value int, pos int)
	RemoveLast()
	Remove(pos int)
	Get(pos int) int
	Size() int
}

type ArrayList struct {
	values   []int // array
	inserted int   // quantidade de numeros inseridos no array
}

// Função para duplicar um vetor (função privada)
func (list *ArrayList) doubleArray() {
	doubleList := make([]int, len(list.values)*2)

	for i := 0; i < len(list.values); i++ {
		doubleList[i] = list.values[i]
	}

	list.values = doubleList // atribue a nova lista para a que já estava em uso, para reutilizar a nomenclatura
}

/* - Função para adicionar elementos no vetor (Função pública)
    Adiciona-se da esquerda para direita, após o ultimo elemento inserido anteriormente.

    - Desempenho
    Melhor caso: Ômega(1)
            Adicionar na última posição
    Pior caso: O(n)
            Ter que duplicar o vetor é mais custoso computacionalmente
*/
func (list *ArrayList) Add(value int) {
	// Verificar se o array está lotado, se sim, ampliá-lo
	if list.inserted >= len(list.values) {
		list.doubleArray()
	}

	list.values[list.inserted] = value
	list.inserted++
}

/* - Função para adicionar um elemento em qualquer posição
    Adiciona-se o elemento depois de afastar os demais elementos para a direita dessa posição.

    - Desempenho
    Melhor caso: Ômega(1)
            Adicionar na última posição
    Pior caso: O(n)
            1) Ter que duplicar o vetor, pois há um custo computacionalmente
            2) Ter que deslocar os elementos à direta da posição desejada para se inserir um novo valor
*/
func (list *ArrayList) AddPos(value int, pos int) {
	// Confere se a posição é válida
	if pos>=0 && pos<list.inserted {
		if list.inserted >= len(list.values) {
			list.doubleArray()
		}
        
        for i:=list.inserted; i>pos; i--{
            list.values[i] = list.values[i-1]
        }

        list.values[pos] = value
        list.inserted++
	}
}

// Substitui um valor por outro (Set)
func (list *ArrayList) Update(value int, pos int){
    // Confere se a posição é válida
	if pos>=0 && pos<list.inserted {
        list.values[pos] = value
    }
}

// Função para remover o último elemento do vetor
func (list *ArrayList) RemoveLast(){
    list.inserted--
}


/* -Função para remover um elemento do vetor
    Remove o elemento e desloca os demais elementos à direta do removido para a esquerda, fechando a lacuna feita pela remoção.

    - Desempenho:
    Melhor caso: Ômega(1)
        Remover o ultimo elemento.
    Pior caso: O(n)
        Remover um elemento no início e ter que deslocar os demais à esquerda.
*/
func (list *ArrayList) Remove(pos int){
    // Confere se a posição é válida
	if pos>=0 && pos<list.inserted {
        
        for i:=pos; pos<list.inserted; i++{
            list.values[i] = list.values[i+1]
        }
        
        list.inserted--
    }    
}

// Função para acessar um elemento do vetor
func (list *ArrayList) Get(pos int) int{
    // Confere se a posição é válida
	if pos>=0 && pos<list.inserted {
        return list.values[pos]
    }
    return 0
}

// Função para retornar o tamanho do vetor
func (list *ArrayList) Size() int{
    return len(list.values)
    // ou return list.inserted 
}

// Fazer testes adicionais para implementar gráficos referentes ao tempo de cada função

// Testar 
func main (){

    var list List = &ArrayList{
        values: make([]int, 10),
        inserted : 0,
    }
    
    list.Add(24)
    list.Add(8)
    list.Add(18)
    list.Add(35)
    list.AddPos(31, 2)

    fmt.Print("[ ")
    for i:=0; i<list.Size(); i++ {
        fmt.Print(list.Get(i), " ")
    }
    fmt.Print("]")
    
}