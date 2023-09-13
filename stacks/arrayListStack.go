package main

type Pilhas interface {
    Push (value int) // empilhar
    Pop ()           // desempilhar
    Top () int       // topo, refere-se ao elemento recentemente inserido
    Empty () bool    // verifica se a pilha está vazia
    Size () int      // tamanho
}

type ArrayListStack struct {
	values []int
	tam int
}

func (al_stack *ArrayListStack) double() {
    // criar um novo array com o dobro de tamanho (capacidade)
    double := make([]int, len(al_stack.values)*2)

    // passa os valores para o novo vetor
    for i:=0; i<len(al_stack.values); i++{
        double[i] = al_stack.values[i]
    }

    // vetor original agora é vetor duplicado
    al_stack.values = double
}

// Adiciona na ultima posição do array
func (al_stack *ArrayListStack) Push (value int) {
    if al_stack.tam == len(al_stack.values){
        al_stack.double()
    }
    al_stack.values[al_stack.tam] = value
    al_stack.tam++
}

func (al_stack *ArrayListStack) Pop () {
    al_stack.tam--
}

func (al_stack *ArrayListStack) Top () int {
    return al_stack.values[al_stack.tam-1]
}

func (al_stack *ArrayListStack) Empty () bool {
    return al_stack.tam == 0
}

func (al_stack *ArrayListStack) Size () int {
    return al_stack.tam
}