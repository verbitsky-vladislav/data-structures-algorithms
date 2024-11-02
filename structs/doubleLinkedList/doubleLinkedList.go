package doubleLinkedList

type LRUCache interface {
}

type Node struct {
	prev *Node
	next *Node
	data string // we need string because we will save key
}

type DoubleLinkedList struct {
	head, tail *Node
	length     int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// Len get length of double linked list
func (ll *DoubleLinkedList) Len() int {
	return ll.length
}

// Front get pointer to head node
func (ll *DoubleLinkedList) Front() *Node {
	return ll.head
}

// Back get pointer to tail node
func (ll *DoubleLinkedList) Back() *Node {
	return ll.tail
}

// MoveToFront move existing node to front
func (ll *DoubleLinkedList) MoveToFront(node *Node) *Node {
	// варианты:
	// 1. Просто поменять дату у элемена node и элемента head
	//
	return node
}

//    a           b         c         d
// [nil, b] -> [a, c] -> [b, d] -> [c, nil]
