package structs

type Node struct {
	next *Node
}

func NewNode() Node {
	return Node{}
}

type LinkedList struct {
	nodes []Node
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (l LinkedList) getLength() int {
	return len(l.nodes)
}

func (l *LinkedList) insertAt(index int, node Node) error {
	// if index > len(l.nodes)-1 {
	// 	return errors.New("Index doesn't exist")
	// }
	// if len(l.nodes) == 0 {
	// 	l.nodes[0] = node
	// 	return nil
	// }
	//
	// before := l.nodes[index-1]
	// l.nodes[index].next = before.next
	// before.next = &new
	return nil
}

func (l *LinkedList) append(node Node) {
	// 0,0
	l.nodes[len(l.nodes)] = node
	l.nodes[len(l.nodes)-1].next = &node
}

func (l *LinkedList) prepend(node Node) {
	l.nodes[0] = node
}

type List interface {
	getLength() int
	insertAt(index int, node Node) error
	append(node Node)
	prepend(node Node)
	delete(node Node) (Node error)
	deleteAt(index int) (Node error)
	get(index int) (Node error)
}
