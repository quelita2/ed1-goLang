package main

type Pilhas interface {
    Push (value int) // empilhar
    Pop ()           // desempilhar
    Top () int       // topo, refere-se ao elemento recentemente inserido
    Empty () bool    // verifica se a pilha está vazia
    Size () int      // tamanho
}

type Node struct {
    value int
    next *Node
}

type LinkedListStack struct {
    tail *Node // último elemento inserido é o topo da pilha
    inserted int
}

// o ultimo aponta para o penultimo...
func (stack *LinkedListStack) Push (value int) {
	// o novo nó aponta para o topo, e o novo nó é o novo topo
    stack.tail = &Node{value, stack.tail}
    stack.inserted ++
}

// remoção do ultimo empilhado, que é o primeiro da pilha
func (stack *LinkedListStack) Pop () {
    // verificar se a pilha tem algo para desempilhar
    if stack.inserted > 0 {
        if stack.tail.next!=nil{
            stack.tail = stack.tail.next
        }else {
            stack.tail = nil
        }
        stack.inserted --
    }
}

func (stack *LinkedListStack) Top () int {
    // se não tiver nada ele retornará nil?
    return stack.tail.value
}

func (stack *LinkedListStack) Empty () bool {
    return stack.inserted == 0
}

func (stack *LinkedListStack) Size () int {
    return stack.inserted
}