package main

type List interface {
	Add(value int)
	AddPos(value int, pos int)
	RemoveLast()
	RemovePos(pos int)
	Get(pos int) int
	Set(value int, pos int)
	Size() int
}

type Node struct {
  value int
  next *Node
}

type LinkedList struct {
	head *Node
	inserted  int
}

func (list *LinkedList) Add(val int){
    newNode := Node{val, nil}

    /* Adicionar o valor novo num nó, verificar se o head já está com valor. Caso não esteja com nenhum valor, aplicar o novo nó. Caso esteja inserido, fazer o ultimo nó apontar para o mais novo nó. Para isso devemos percorrer a lista.
    */
    if list.head == nil{
        list.head = &newNode // aponta para o endereço do novo nó
    } else {
        // Percorrendo a lista para encontrar o ultimo nó
        aux := list.head
        for aux.next != nil{
            aux = aux.next
        }
        aux.next = &newNode
    }
    list.inserted ++
}

func (list *LinkedList) AddPos (val int, pos int){
    // Verifica se a posição é válida
    if pos>=0 && pos<=list.inserted{
        newNode := Node{val, nil}
        
        // Adicionar no inicio, na posição 0:
        if pos == 0 {
            newNode.next = list.head // que agora será o 2° elemento
            list.head = &newNode // atualizando o 1° elemento
        } else {
            
            // Percorrer a lista até a posição válida com variáveis auxiliares
            aux := list.head
            for i:=1; aux.next != nil && i<pos; i++ { // contador a partir de 1 para parar antes da posição desejada
                aux = aux.next
            }
            newNode.next = aux.next
            aux.next = &newNode
        }
        list.inserted ++
    }
}

func (list *LinkedList) RemoveLast(){
    // para remover o ultimo será preciso percorrer até o penultimo elemento e aponta-lo para nulo ou, em caso de só haver 1 elemento na lista, fazer o list.head virar nulo. Logo depois, decrementamos o numero de inseridos
    if list.head != nil{
        if list.inserted == 1{
            list.head = nil
        }else{
            aux := list.head
            aux_ant := aux
            for aux.next!=nil {
                aux_ant = aux
                aux = aux.next
            }
            aux_ant.next = nil
        }
        list.inserted --
    }
}

func (list *LinkedList) RemovePos(pos int){
    // Verificar se a posição é válida
    if pos>=0 && pos<=list.inserted{
        
        if pos == 0{ // remoção da posição inicial
            list.head = list.head.next
        } else { // remoção de um elemento no meio ou no final da lista
            // Percorrer com um contador até a posição solicitada
            aux := list.head
            aux_ant := aux
            for i:=0; i<pos; i++ {
                aux_ant = aux
                aux = aux.next
            }
            aux_ant.next = aux.next
            aux.next = nil
        }
        list.inserted --
    }
}

func (list *LinkedList) Get(pos int) int{
    if pos >=0 && pos <= list.inserted {
        if pos == 0{
            return list.head.value
        } else {
            aux:= list.head
            for i:=0; i<pos; i++ {
                aux = aux.next
            }
            return aux.value
        }
    }
    return -1
}

func (list *LinkedList) Set(val int, pos int){
    if pos>=0 && pos<=list.inserted {
        if pos == 0 {
            list.head.value = val
        } else {
            aux:=list.head
            for i:=0; i<pos; i++ {
                aux = aux.next
            }
            aux.value = val
        }
    }
}

func (list *LinkedList) Size() int{
    return list.inserted
}