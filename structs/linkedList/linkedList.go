package linkedList

import (
	"fmt"
)

// LinkedListInterface определяет методы для работы с односвязным списком
// Индекс головы начинается с 1, но это можно изменить
type LinkedListInterface interface {
	// Traverse - Обходит все узлы в списке и выводит их значения
	// Время выполнения: O(n), Память: O(1), так как не требуется дополнительное хранение данных
	Traverse()

	// Search - Ищет узел с указанным значением и возвращает указатель на этот узел
	// Время выполнения: O(n), Память: O(1), так как используется только указатель на текущий узел
	Search(value interface{}) *Node

	// Length - Возвращает длину списка
	// Время выполнения: O(n), Память: O(1), так как используется только счетчик для подсчета узлов
	Length() int

	// InsertFront - Вставляет новый узел в начало списка
	// Время выполнения: O(1), Память: O(1), так как создается один новый узел
	InsertFront(data interface{}) *Node

	// InsertBack - Вставляет новый узел в конец списка
	// Время выполнения: O(1), Память: O(1), так как создается один новый узел
	InsertBack(data interface{}) *Node

	// InsertIntoPosition - Вставляет новый узел на указанную позицию в списке
	// Время выполнения: O(n), Память: O(1), поскольку создается один новый узел
	InsertIntoPosition(data interface{}, position int) *Node

	// DeleteFront - Удаляет первый узел из списка
	// Время выполнения: O(1), Память: O(1), так как не требуется дополнительное хранение данных
	DeleteFront()

	// DeleteBack - Удаляет последний узел из списка
	// Время выполнения: O(n), Память: O(1), так как используется указатель для прохода по списку
	DeleteBack()

	// DeleteFromPosition - Удаляет узел на указанной позиции в списке
	// Время выполнения: O(n), Память: O(1), так как используется указатель для прохода по списку
	DeleteFromPosition(position int)

	// Update - Обновляет данные узла на указанной позиции
	// Время выполнения: O(n), Память: O(1), так как изменяется только значение узла
	Update(position int, data interface{})

	// IsEmpty - Проверяет, пуст ли список
	// Время выполнения: O(1), Память: O(1), так как проверяется только указатель на голову списка
	IsEmpty() bool

	// Clear - Очищает список, удаляя все узлы
	// Время выполнения: O(1), Память: O(1), так как сбрасываются только указатели на голову и хвост
	Clear()

	// GetAtPosition - Возвращает данные узла на указанной позиции
	// Время выполнения: O(n), Память: O(1), так как используется указатель для прохода по списку
	GetAtPosition(position int) (interface{}, error)

	// Reverse - Разворачивает список, меняя порядок узлов на противоположный
	// Итеративный метод: Время выполнения: O(n), Память: O(1), так как указатели меняются на месте
	// Рекурсивный метод: Время выполнения: O(n), Память: O(n), так как каждый рекурсивный вызов требует дополнительного места в стеке
	Reverse()

	// ReverseTraversal - Выполняет обход списка в обратном порядке с использованием рекурсии
	// Время выполнения: O(n), Память: O(n), поскольку каждый вызов рекурсии требует дополнительного места в стеке
	ReverseTraversal()

	// Merge - Сливает другой список с текущим
	// Время выполнения: O(n), Память: O(1) для простого случая объединения списков
	Merge(other *LinkedList) // todo пока ленб делать - потом сделатб надо

	// DeleteByValue - Удаляет первый узел с указанным значением
	// Время выполнения: O(n), Память: O(1), так как требуется указатель для поиска и удаления узла
	DeleteByValue(value interface{}) // todo пока ленб делатб - потом сделать надо
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

// ReverseTraversal using recursion – O(n) Time and O(n) Space
func (ll *LinkedList) ReverseTraversal() {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return
	}
	ll.reverseTraversalRecursive(ll.head)
}

// helper function because we need params
func (ll *LinkedList) reverseTraversalRecursive(node *Node) {
	if node == nil {
		return
	}
	ll.reverseTraversalRecursive(node.next)

	fmt.Printf("Node: %v\n", node.data)
}

// Reverse
// Вход : Связанный список = 1 -> 2 -> 3 -> 4 -> NULL
// Выход : Обратный связанный список = 4 -> 3 -> 2 -> 1 -> NULL
func (ll *LinkedList) Reverse() {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return
	}

	var previous, next *Node
	current := ll.head

	// меняем хвост и голову
	ll.tail = ll.head

	for current != nil {
		next = current.next
		current.next = previous

		previous = current
		current = next
	}

	ll.head = previous
}

func (ll *LinkedList) Search(value interface{}) *Node {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}

	currentNode := ll.head
	for currentNode != nil {
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
	for currentNode != nil {
		length = length + 1
		currentNode = currentNode.next
	}
	return length
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.head == nil && ll.tail == nil
}

func (ll *LinkedList) InsertFront(data interface{}) *Node {
	newNode := &Node{
		data: data,
		next: ll.head,
	}
	ll.head = newNode

	if ll.tail == nil {
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
		return nil
	}

	newNode := Node{
		data: data,
	}
	// если position == 1, узел вставляется в начало, и новый узел становится головой списка
	if position == 1 {
		newNode.next = ll.head
		ll.head = &newNode
		if ll.tail == nil {
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

	if newNode.next == nil {
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
	if ll.head == nil {
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
		ll.tail = nil
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
		if ll.head == nil {
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

	if currentNode.next == nil {
		ll.tail = currentNode
	}
}

func (ll *LinkedList) Update(position int, data interface{}) {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return
	}

	// если position < 1, возвращаем ошибку
	if position < 1 {
		fmt.Printf("Invalid position %v", position)
		return
	}

	// если position == 0, то изменяемый узел - голова
	if position == 1 {
		ll.head.data = data
		return
	}

	currentNode := ll.head

	// дополнительные проверки в цикле + смотрим что мы не выходим за пределы списка
	for i := 0; i < position && currentNode != nil; i++ {
		currentNode = currentNode.next
	}

	// если currentNode == nil или currentNode.next == nil, позиция выходит за пределы списка
	if currentNode == nil {
		fmt.Printf("Position: %v out of bounds\n", position)
		return
	}

	currentNode.data = data

	return

}

func (ll *LinkedList) Clear() {
	ll.head = nil
	ll.tail = nil
}

func (ll *LinkedList) GetAtPosition(position int) (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("linked list is empty")
	}
	if position < 1 {
		return nil, fmt.Errorf("invalid position %v", position)
	}
	if position == 1 {
		return ll.head.data, nil
	}

	currentNode := ll.head

	for i := 0; currentNode != nil && i < position; i++ {
		currentNode = currentNode.next
	}

	// вышли за пределы списка -> кидаем ошибку
	if currentNode == nil {
		return nil, fmt.Errorf("position: %v out of bounds", position)
	}

	return currentNode.data, nil

}

func (ll *LinkedList) Merge(other *LinkedList) {

}

func (ll *LinkedList) DeleteByValue(value interface{}) {

}
