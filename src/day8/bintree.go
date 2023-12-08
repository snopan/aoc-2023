package day8

type Node struct {
	value string
	left  *Node
	right *Node
}

func NewNode(value string) *Node {
	return &Node{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (n *Node) SetLeftChildren(node *Node) {
	n.left = node
}

func (n *Node) SetRightChildren(node *Node) {
	n.right = node
}

func (n *Node) AddLeftChildren(value string) *Node {
	return n.addChildren(true, value)
}

func (n *Node) AddRightChildren(value string) *Node {
	return n.addChildren(false, value)
}

func (n *Node) addChildren(left bool, value string) *Node {
	node := NewNode(value)
	returnNode := node
	if value == n.value {
		node = n
		returnNode = nil
	}

	if left {
		n.SetLeftChildren(node)
	} else {
		n.SetRightChildren(node)
	}

	return returnNode
}
