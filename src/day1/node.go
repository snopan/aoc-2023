package day1

type Node struct {
	value    int
	children map[string]*Node
}

func (n *Node) addNode(letter string, value int) *Node {
	child, ok := n.children[letter]
	if !ok {
		child = &Node{
			value:    value,
			children: map[string]*Node{},
		}
		n.children[letter] = child
	}
	return child
}

func (n *Node) validValue() bool {
	return n.value != -1
}
