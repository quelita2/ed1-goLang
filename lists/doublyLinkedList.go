package main

type List interface {
	Add(value int)
	AddAtIndex(index int, value int)
	Remove(index int)
	Get(index int) int
	Set(index int, value int)
	Size() int
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	inserted int
}

type Node struct {
	prev *Node // anterior
	value int
	next *Node // próximo
}

func (list *DoublyLinkedList) Add(value int) {
	newNode := Node{nil, value, nil}

	if list.head == nil {
		// add o nov nó como sendo o primeiro e ultimo, já que é o unico
		list.head = &newNode
		list.tail = &newNode
	} else {
		// direciona até o último elemento e aponta para o novo nó, enquanto que o novo nó aponta para o novo penúltimo nó
		aux := list.tail
		aux.next = &newNode
		newNode.prev = aux
	}
	list.inserted++
}

func (list *DoublyLinkedList) AddAtIndex(index int, value int) {
	// verifica se o indice é valido
	if index >= 0 && index <= list.inserted {
		newNode := Node{nil, value, nil}

		// indice == 0
		if index == 0 {
			if list.head != nil {
				newNode.next = list.head
			}
			list.head = &newNode
		} else {
			// indice no meio/final da lista duplamente ligada
			aux := list.head
			for i := 0; i < index-1; i++ { // varre a lista até a cauda
				aux = aux.next
			}
			// ligando da esquerda para direita
			newNode.next = aux.next
			newNode.prev = aux
			// ligando da direita para esquerda
			aux.next = &newNode
			if aux.next != nil {
				newNode.next.prev = &newNode
			}
		}
		list.inserted++
	}
}

func (list *DoublyLinkedList) Remove(index int) {
	if index >= 0 && index <= list.inserted {
		if index == 0 {
			if list.head != nil {
				list.head = list.head.next
				list.head.prev = nil
				if list.inserted == 1 {
					list.tail = nil
				}
			}
		} else {
			aux := list.head
			aux_ant := list.head
			for i := 0; i < index; i++ {
				aux_ant = aux
				aux = aux.next
			}
            aux_ant.next = aux.next
            if aux.next != nil {
                aux.next.prev = aux_ant
            }
		}
	}
	list.inserted--
}

func (list *DoublyLinkedList) Get(index int) int {
    if index>=0 && index<=list.inserted {
        if index == 0 {
            return list.head.value
        } else {
            aux := list.head
            for i:=0; i<index; i++ {
                aux = aux.next
            }
            return aux.next.value
        }
    }
    return -1
}

func (list *DoublyLinkedList) Set(index int, value int) {
    if index>=0 && index<=list.inserted {
        if index == 0 {
            list.head.value = value
        } else {
            aux := list.head
            for i:=0; i<index; i++ {
                aux = aux.next
            }
            aux.value = value
        }
    }
}