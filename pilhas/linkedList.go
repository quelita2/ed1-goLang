package pilhas

import ("errors")

type LinkedList_Stack struct{
    head *Node
    tam int
}

type Node struct{
    value int
    next *Node
}

// Alocar e incializar os valores da LinkedList
func (ll_stack *LinkedList_Stack) Init(){
    ll_stack.head = nil
    ll_stack.tam = 0
}

// Inserir um elemento no topo da pilha: Push()
func (ll_stack *LinkedList_Stack) Push(value int){
    ll_stack.head = &Node{value: value, next: ll_stack.head}
    ll_stack.tam++
}

// Remover um elemento do topo da pilha: Pop()
func (ll_stack *LinkedList_Stack) Pop() (int, error){
    //erro: querer tirar mais do que a linkedlist possue
    if ll_stack.tam > 0{
        ll_stack.head = ll_stack.head.next
        ll_stack.tam--
    }else{
        return -1, errors.New("Não há elementos para se remover da pilha")
    }
	return -1, nil
}

// Ler o elemento do topo
func (ll_stack *LinkedList_Stack) Top() (int, error){
    if ll_stack.tam > 0{
        return ll_stack.head.value, nil
    }else{
		return -1, errors.New("Não há elementos no topo")
	}
}

// Contar os elementos de uma pilha
func (ll_stack *LinkedList_Stack) Size() int{
    return ll_stack.tam
}

// Perguntar se a pilha está vazia
func (ll_stack *LinkedList_Stack) Empty() bool{
    if ll_stack.tam != 0{
        return false
    }else{
        return true
    }
}