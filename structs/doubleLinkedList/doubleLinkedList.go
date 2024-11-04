package doubleLinkedList

import (
	"fmt"
)

type DoublyLinkedListInterface interface {
	// Traverse - Обходит все узлы в прямом порядке и выводит их значения
	// Время выполнения: O(n), Память: O(1), так как не требуется дополнительное хранение данных
	Traverse()

	// ReverseTraverse - Обходит все узлы в обратном порядке, используя указатель на предыдущий узел
	// Время выполнения: O(n), Память: O(1), так как обход выполняется через связи на предыдущие узлы
	ReverseTraverse()

	// Search - Ищет узел с указанным значением и возвращает указатель на этот узел
	// Время выполнения: O(n), Память: O(1)
	Search(value interface{}) *Node

	// Length - Возвращает длину списка
	// Время выполнения: O(n), Память: O(1)
	Length() int

	// InsertFront - Вставляет новый узел в начало списка
	// Время выполнения: O(1), Память: O(1)
	InsertFront(data interface{}) *Node

	// InsertBack - Вставляет новый узел в конец списка
	// Время выполнения: O(1), Память: O(1)
	InsertBack(data interface{}) *Node

	// InsertIntoPosition - Вставляет новый узел на указанную позицию в списке
	// Время выполнения: O(n), Память: O(1)
	InsertIntoPosition(data interface{}, position int) *Node

	// DeleteFront - Удаляет первый узел из списка
	// Время выполнения: O(1), Память: O(1)
	DeleteFront()

	// DeleteBack - Удаляет последний узел из списка, используя указатель на предыдущий узел
	// Время выполнения: O(1), Память: O(1)
	DeleteBack()

	// DeleteFromPosition - Удаляет узел на указанной позиции в списке
	// Время выполнения: O(n), Память: O(1)
	DeleteFromPosition(position int)

	// DeleteByValue - Удаляет первый узел с указанным значением
	// Время выполнения: O(n), Память: O(1)
	DeleteByValue(value interface{})

	// Update - Обновляет данные узла на указанной позиции
	// Время выполнения: O(n), Память: O(1)
	Update(position int, data interface{})

	// IsEmpty - Проверяет, пуст ли список
	// Время выполнения: O(1), Память: O(1)
	IsEmpty() bool

	// Clear - Очищает список, удаляя все узлы
	// Время выполнения: O(1), Память: O(1)
	Clear()

	// GetAtPosition - Возвращает данные узла на указанной позиции
	// Время выполнения: O(n), Память: O(1)
	GetAtPosition(position int) (interface{}, error)

	// Reverse - Разворачивает список, меняя порядок узлов на противоположный
	// Итеративный метод: Время выполнения: O(n), Память: O(1)
	Reverse()

	// Merge - Сливает другой двусвязный список с текущим
	// Время выполнения: O(n), Память: O(1)
	Merge(other *DoubleLinkedList)
}

type Node struct {
	prev *Node
	next *Node
	data interface{}
}

type DoubleLinkedList struct {
	head, tail *Node
}

func (dll *DoubleLinkedList) Traverse() {
	if dll.head == nil {
		fmt.Println("Linked list is empty.")
		return
	}

	position := 0
	current := dll.head

	fmt.Println("Starting linked list traversal:")
	for current != nil {
		fmt.Printf("Node %d: %v\n", position, current.data)

		current = current.next
		position++
	}
}

func (dll *DoubleLinkedList) ReverseTraverse() {
	if dll.head == nil {
		fmt.Println("Linked list is empty.")
		return
	}
	fmt.Println("Starting linked list traversal:")
	// если есть голова и нет конца, то мы просто логнем голову и выйдем
	if dll.tail == nil {
		fmt.Printf("Head node: %v\n", dll.head)
		return
	}

	current := dll.tail
	for current != nil {
		fmt.Printf("Node: %v\n", current.data)
		current = current.prev
	}
}

func (dll *DoubleLinkedList) Search(value interface{}) *Node {
	if dll.head == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}

	current := dll.head
	for current != nil {
		if current.data == value {
			return current
		}
		current = current.next
	}
	return nil
}

func (dll *DoubleLinkedList) Length() int {
	if dll.head == nil {
		fmt.Println("Linked list is empty.")
		return 0
	}
	length := 0
	current := dll.head
	for current != nil {
		length++
		current = current.next
	}

	return length
}

func (dll *DoubleLinkedList) InsertFront(data interface{}) *Node {
	newNode := &Node{
		data: data,
		next: nil,
		prev: nil,
	}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return newNode
	}

	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode

	return newNode
}

func (dll *DoubleLinkedList) InsertBack(data interface{}) *Node {
	newNode := &Node{
		data: data,
		next: nil,
		prev: nil,
	}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return newNode
	}

	dll.tail.next = newNode
	newNode.prev = dll.tail
	dll.tail = newNode

	return newNode
}

func (dll *DoubleLinkedList) InsertAtPosition(data interface{}, position int) *Node {
	newNode := &Node{
		data: data,
		next: nil,
		prev: nil,
	}

	if position < 0 {
		fmt.Printf("Invalid position %v", position)
		return nil
	}

	if position == 1 {
		newNode.next = dll.head
		if dll.head != nil {
			dll.head.prev = newNode
		}
		dll.head = newNode

		if dll.tail == nil {
			dll.tail = newNode
		}
		return newNode
	}

	current := dll.head
	for i := 1; i < position-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		fmt.Printf("Position: %v out of bounds", position)
		return nil
	}

	newNode.prev = current
	newNode.next = current.next

	if current.next != nil {
		current.next.prev = newNode
	}
	current.next = newNode

	if newNode.next == nil {
		dll.tail = newNode
	}

	return newNode
}
