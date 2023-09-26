package main

import (
	"errors"
)

type IQueue interface {
	Enqueue(value int)     // Enfileirar
	Dequeue() (int, error) // Desenfileirar
	Front() (int, error)   // Retorna o primeiro da fila
	IsEmpty() bool         // Retorna se a fila está vazia
	Size() int             // Retorna o tamanho da fila
}

type LinkedListQueue struct {
	front    *Node
	rear     *Node
	inserted int
}

type Node struct {
	value int
	next  *Node
}

func (queue *LinkedListQueue) Enqueue(value int) {
	newNode := Node{value, nil}

	if queue.inserted == 0 {
		queue.front = &newNode // aponta para o 1° da fila
	} else {
		queue.rear.next = &newNode // o ultimo aponta para o novo ultimo
	}
	queue.rear = &newNode // atualiza o ultimo da fila
	queue.inserted++
}

func (queue *LinkedListQueue) Dequeue() (int, error) { // retorna o valor que saiu da fila
	if queue.inserted == 0 {
		return -1, errors.New("Não é possível desenfileirar, pois já está vazia!")
	} else {
		lastfront := queue.front.value
		if queue.inserted == 1 {
			queue.front = nil
			queue.rear = nil
		} else {
			queue.front = queue.front.next
		}
		queue.inserted--
		return lastfront, nil
	}
}

func (queue *LinkedListQueue) Front() (int, error) {
	if queue.front.value == 0 {
		return -1, errors.New("A fila está vazia")
	} else {
		return queue.front.value, nil
	}
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.inserted == 0
}

func (queue *LinkedListQueue) Size() int {
	return queue.inserted
}
