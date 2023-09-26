package main

import "errors"

type IQueue interface {
	Enqueue(value int)
	Dequeue() (int, error)
	Front() (int, error)
	IsEmpty() bool
	Size() int
}

type ArrayQueue struct {
    v []int
    front int
    rear int
    inserted int
    capacity int // capacidade do array
}

func (queue *ArrayQueue) Enqueue(value int) {
    if queue.inserted == 0 {
        queue.front = 0
        queue.rear = 0
        queue.v[0] = value
    } else {
        queue.rear = (queue.rear+1) % queue.capacity
        queue.v[queue.rear] = value
    }
    queue.inserted++
}

func (queue *ArrayQueue) Dequeue() (int, error) {
    if queue.inserted == 0 {
        return -1, errors.New("Não há oq desenfileirar, a fila está vazia!")
    } else {
        aux := queue.v[queue.front]
        
        if queue.inserted == 1{
            queue.front = -1
            queue.rear = -1
        }
        
        queue.front = (queue.front+1) % queue.capacity
        queue.inserted--
        
        return aux, nil
    }
}

func (queue *ArrayQueue) Front() (int, error) {
    if queue.front <= 0 {
        return -1, errors.New("A fila está vazia para desenfileirar!")
    } else {
        return queue.front, nil
    }
}

func (queue *ArrayQueue) IsEmpty() bool {
    return queue.inserted == 0
}

func (queue *ArrayQueue) Size() int {
    return queue.inserted
}