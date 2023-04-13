package pilhas

import (
	"errors"
)

type ArrayListStack struct {
	values []int
	tam    int
}

// Alocar e incializar os valores do ArrayList
func (arrayliststack *ArrayListStack) Init() {
	arrayliststack.tam = 0
	arrayliststack.values = make([]int, arrayliststack.tam)
}

// Duplica o arrayliststack para compactar mais valores
func (arrayliststack *ArrayListStack) double() {
	doubledValues := make([]int, len(arrayliststack.values)*2)

	for i := 0; i < arrayliststack.tam; i++ {
		doubledValues[i] = arrayliststack.values[i]
	}

	arrayliststack.values = doubledValues
}

// Inserir um elemento no topo da pilha
func (arrayliststack *ArrayListStack) Push(value int) {
	if arrayliststack.tam == len(arrayliststack.values) {
		arrayliststack.double()
	}
	arrayliststack.values[arrayliststack.tam] = value
	arrayliststack.tam++
}

// Remover um elemento do topo da pilha
func (arrayliststack *ArrayListStack) Pop() error {
	if arrayliststack.tam == 0 {
		return errors.New("Não há posições para excluir")
	}

	if arrayliststack.tam > 0 {
		arrayliststack.tam--
	}
	return nil
}

// Ler o elemento do topo
func (arrayliststack *ArrayListStack) Top() int{
  return arrayliststack.values[arrayliststack.tam]
}

// Contar os elementos de uma pilha
func (arrayliststack *ArrayListStack) Size() int{
  return len(arrayliststack.values)
}

// Perguntar se a pilha está vazia
func (arrayliststack *ArrayListStack) Empty() bool{
  if arrayliststack.tam != 0{
    return false
  }else{
    return true
  }
}