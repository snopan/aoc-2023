package day1

type node struct {
	value    int
	children map[string]*node
}

func (n *node) addNode(letter string, value int) *node {
	child, ok := n.children[letter]
	if !ok {
		child = &node{
			value:    value,
			children: map[string]*node{},
		}
		n.children[letter] = child
	}
	return child
}

func (n *node) validValue() bool {
	return n.value != -1
}
