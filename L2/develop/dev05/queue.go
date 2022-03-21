package main

import (
	"errors"
)

//Queue - очередь для строковых элементов. Length - указывает сколько сейчас элементов в очереди, а Size сколько может хранить очередь.
type Queue struct {
	Length  int
	Size    int
	FirstEl *ElementQueue
	LastEl  *ElementQueue
}

//NewQueue - создает новую очередь с указным размером.
func NewQueue(size int) *Queue {
	return &Queue{
		Length: 0,
		Size:   size,
	}
}

//Enqueue - добавляет элемент в очередь.
func (q *Queue) Enqueue(v string) {
	if q.Length == q.Size {
		q.FirstEl = q.FirstEl.Next
	} else {
		q.Length++
	}
	
	if !q.Empty() {
		q.FirstEl=NewElementQueue(v)
		q.LastEl=NewElementQueue(v)
		return
	}

	q.LastEl.Next = NewElementQueue(v)
	q.LastEl = q.LastEl.Next
}

//Dequeue - выдает элемент из очереди и удаляет его.
func (q *Queue) Dequeue() (string, error) {
	if q.Empty() {
		return "", errors.New("нет элементов")
	}

	str := q.FirstEl.Value
	if q.Empty() {
		q.FirstEl = q.FirstEl.Next
	} else {
		q.FirstEl = nil
	}
	q.Length--
	return str, nil
}

func (q *Queue) Show() (string, error) {
	if q.Empty() {
		return "", errors.New("нет элементов")
	}

	str := q.FirstEl.Value
	if q.Empty() {
		q.FirstEl = q.FirstEl.Next
	} else {
		q.FirstEl = nil
	}
	q.Length--
	return str, nil
}

//Empty - проверяет есть ли элементы в очереди.
func (q *Queue) Empty() bool {
	return q.Length > 0
}

//ElementQueue - элемент очереди.
type ElementQueue struct {
	Next  *ElementQueue
	Value string
}

//NewElementQueue - создает указатель на новый элемент очереди с заданым значением.
func NewElementQueue(v string) *ElementQueue {
	return &ElementQueue{
		Value: v,
	}
}
