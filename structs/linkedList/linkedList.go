package linkedList

import (
	"fmt"
)

type LinkedListInterface interface {
	// Traverse - проходит по всем узлам списка и выводит их значения
	Traverse()
	// Searching - ищет и возвращает указатель на узел с указанным значением
	Searching(value interface{}) *Node
	// Length - возвращает длину списка
	Length() int
	// InsertFront - вставляет новый узел в начало списка
	InsertFront(data interface{}) *Node
	// InsertBack - вставляет новый узел в конец списка
	InsertBack(data interface{}) *Node
	// InsertIntoPosition - вставляет новый узел в заданную позицию списка
	InsertIntoPosition(data interface{}, position int) *Node
	// DeleteFront - удаляет первый узел из списка
	DeleteFront()
	// DeleteBack - удаляет последний узел из списка
	DeleteBack()
	// DeleteFromPosition - удаляет узел в указанной позиции
	DeleteFromPosition(position int)

	// Update - обновляет данные узла в заданной позиции
	Update(position int, data interface{}) // todo
	// IsEmpty - проверяет, пуст ли список
	IsEmpty() bool // todo
	// Clear - удаляет все узлы из списка, очищая его
	Clear() // todo
	// GetAtPosition - возвращает данные узла на заданной позиции
	GetAtPosition(position int) (interface{}, error) // todo
	// ReverseTraversal - выполняет обход списка в обратном порядке
	ReverseTraversal() // todo
	// Reverse - разворачивает список, меняя порядок узлов на противоположный
	Reverse() // todo
	// Merge - сливает другой список с текущим
	Merge(other *LinkedList) // todo
	// DeleteByValue - удаляет первый узел с указанным значением
	DeleteByValue(value interface{}) // todo
}

type Node struct {
	next *Node
	data interface{}
}

type LinkedList struct {
	head, tail *Node
}

func New() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
	}
}

func (ll *LinkedList) Traverse() {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return
	}

	fmt.Println("Starting linked list traversal:")
	currentNode := ll.head
	position := 1

	for currentNode != nil {
		fmt.Printf("Node %d: %v\n", position, currentNode.data)
		currentNode = currentNode.next
		position++
	}

	fmt.Println("LinkedList traversal complete.")
}

func (ll *LinkedList) Searching(value interface{}) *Node {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}

	currentNode := ll.head
	for currentNode != nil { // todo разобраться почему так
		if currentNode.data == value {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

func (ll *LinkedList) Length() int {
	length := 0
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return length
	}

	currentNode := ll.head
	for currentNode != nil { // todo разобраться почему так
		length = length + 1
		currentNode = currentNode.next
	}
	return length
}

func (ll *LinkedList) InsertFront(data interface{}) *Node {
	newNode := &Node{
		data: data,
		next: ll.head,
	}
	ll.head = newNode

	if ll.tail == nil { // todo разобраться почему так
		ll.tail = newNode
	}

	return newNode
}

func (ll *LinkedList) InsertBack(data interface{}) *Node {
	newNode := Node{
		next: nil,
		data: data,
	}

	// если список пустой, устанавливаем новый узел как голову и хвост
	if ll.head == nil {
		ll.head = &newNode
		ll.tail = &newNode
		return &newNode
	}

	ll.tail.next = &newNode
	ll.tail = &newNode

	return ll.tail
}

func (ll *LinkedList) InsertIntoPosition(data interface{}, position int) *Node {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}
	// если position < 1, возвращаем ошибку
	if position < 1 {
		fmt.Printf("Invalid position %v", position)
		return ll.head
	}

	newNode := Node{
		data: data,
	}
	// если position == 1, узел вставляется в начало, и новый узел становится головой списка
	if position == 1 {
		newNode.next = ll.head
		ll.head = &newNode
		if ll.tail == nil { // todo разобраться почему так
			ll.tail = &newNode
		}
		return &newNode
	}

	currentNode := ll.head

	// дополнительные проверки в цикле + смотрим что мы не выходим за пределы списка
	for i := 0; i < position-1 && currentNode != nil; i++ {
		currentNode = currentNode.next
	}

	// вышли за пределы списка -> кидаем ошибку
	if currentNode == nil {
		fmt.Printf("Position: %v out of bounds", position)
		return nil
	}

	newNode.next = currentNode.next
	currentNode.next = &newNode

	if newNode.next == nil { // todo разобраться почему так
		ll.tail = &newNode
	}

	return &newNode
}

func (ll *LinkedList) DeleteFront() {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return
	}

	ll.head = ll.head.next
	if ll.head == nil { // todo разобраться почему так
		ll.tail = nil
	}
}

func (ll *LinkedList) DeleteBack() {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return
	}

	if ll.head.next == nil {
		ll.head = nil
		ll.tail = nil // todo разобраться почему так
		return
	}

	secondLastNode := ll.head

	for secondLastNode.next.next != nil {
		secondLastNode = secondLastNode.next
	}

	ll.tail = secondLastNode
	ll.tail.next = nil
}

func (ll *LinkedList) DeleteFromPosition(position int) {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return
	}

	// если position < 1, возвращаем ошибку
	if position < 1 {
		fmt.Printf("Invalid position %v", position)
		return
	}

	// если позиция 1, то удаляем голову
	if position == 1 {
		ll.head = ll.head.next
		if ll.head == nil { // todo разобраться почему так
			ll.tail = nil
		}
		return
	}

	currentNode := ll.head
	// дополнительные проверки в цикле + смотрим что мы не выходим за пределы списка
	for i := 0; i < position-2 && currentNode != nil; i++ {
		currentNode = currentNode.next
	}

	// вышли за пределы списка -> кидаем ошибку
	if currentNode == nil || currentNode.next == nil {
		fmt.Printf("Position: %v out of bounds", position)
		return
	}

	currentNode.next = currentNode.next.next

	if currentNode.next == nil { // todo разобраться почему так
		ll.tail = currentNode
	}
}
