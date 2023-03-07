package main

import "fmt"

type Node struct {
	value    string
	linkNode *Node
}

func (n *Node) getValue() string {
	return n.value
}

func (n *Node) getLinkNode() *Node {
	return n.linkNode
}

func (n *Node) setLinkNode(newLinkNode *Node) {
	n.linkNode = newLinkNode
}

func main() {
	yacko := Node{value: "likes to yack"}
	wacko := Node{value: "has a penchant for hoarding snacks"}
	dot := Node{value: "enjoys spending time in movie lots"}

	fmt.Println(yacko)
	fmt.Println(wacko)
	fmt.Println(dot)

	dot.setLinkNode(&wacko)
	yacko.setLinkNode(&dot)

	dotsData := yacko.getLinkNode().getValue()
	wackosData := dot.getLinkNode().getValue()

	fmt.Println(dotsData)
	fmt.Println(wackosData)
}
