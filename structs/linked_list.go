package structs

type List interface {
	// getLength() int
	// insertAt(index int, item Item) error
	// append(item Item)
	// prepend(item Item)
	// delete(item Item) (Item error)
	// deleteAt(index int) (Item error)
	// get(index int) (item error)
}

type Node struct {
	data interface{}
	next *Node
}

func NewNode() Node {
	return Node{}
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (l *LinkedList) append() {

}
