package main

import "errors"

type IDeque interface {
	EnqueueFront(value int)
	EnqueueRear(value int)
	DequeueFront() (int, error)
	DequeueRear() (int, error)
	Front() (int, error)
	Rear() (int, error)
	IsEmpty() bool
	Size() int
}

type DoublyLinkedListDeque struct {
	front *Node
	rear  *Node
	inserted  int
}

type Node struct {
	prev *Node
	value  int
	next *Node
}

func (deque *DoublyLinkedListDeque) EnqueueFront(value int){
    newNode := Node{nil, value, nil}

    if deque.inserted == 0 {
        deque.rear = &newNode
    } else {
        newNode.next = deque.front
        deque.front.prev = &newNode
    }
    deque.front = &newNode
    deque.inserted++
}

func (deque *DoublyLinkedListDeque) EnqueueRear(value int){
    newNode := Node{nil, value, nil}
    
    if deque.inserted == 0 {
        deque.front = &newNode
    } else {
        newNode.prev = deque.rear
        deque.rear.next = &newNode
    }
    deque.rear = &newNode
    deque.inserted ++
}

func (deque *DoublyLinkedListDeque) DequeueFront() (int, error){
    if deque.inserted == 0 {
        return -1, errors.New("A fila duplamente ligada está vazia para se desenfileirar!")
    } else {
        aux := deque.front.value

        if deque.inserted == 1 {
            deque.front = nil
            deque.rear = nil
        } else {
            deque.front = deque.front.next
            deque.front.prev = nil
        }

        deque.inserted --
        return aux, nil
    }
}

func (deque *DoublyLinkedListDeque) DeueueRear() (int, error){
    if deque.inserted == 0 {
        return -1, errors.New("A fila duplamente ligada está vazia para se desenfileirar!")
    } else {
        aux := deque.rear.value

        if deque.inserted == 1 {
            deque.front = nil
            deque.rear = nil
        } else {
            deque.rear.prev.next = nil
            deque.rear = deque.rear.prev
        }
        
        deque.inserted--
        return aux, nil
    }
}

func (deque *DoublyLinkedListDeque) Front() (int, error){
    if deque.inserted <=0 {
        return -1, errors.New("A fila duplamente ligada está vazia!")
    } else {
        return deque.front.value, nil
    }
}

func (deque *DoublyLinkedListDeque) Rear() (int, error){
    if deque.inserted <=0 {
        return -1, errors.New("A fila duplamente ligada está vazia!")
    } else {
        return deque.rear.value, nil
    }
}

func (deque *DoublyLinkedListDeque) IsEmpty() bool{
    return deque.inserted == 0
}

func (deque *DoublyLinkedListDeque) Size() int{
    return deque.inserted
}