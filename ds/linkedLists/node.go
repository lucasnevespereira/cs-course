package main

type Node struct {
	value    string
	nextNode *Node
}

func (n *Node) getValue() string {
	return n.value
}
func (n *Node) getNextNode() *Node {
	return n.nextNode
}
func (n *Node) setNextNode(newNode *Node) {
	n.nextNode = newNode
}
